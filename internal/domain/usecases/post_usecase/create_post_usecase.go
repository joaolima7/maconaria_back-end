package post_usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	post_repository "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/post"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
	"github.com/joaolima7/maconaria_back-end/internal/types"
)

type CreatePostInputDTO struct {
	Title               string         `json:"title" validate:"required"`
	Category            string         `json:"category" validate:"required"`
	SmallDescription    string         `json:"small_description" validate:"required"`
	CompleteDescription string         `json:"complete_description" validate:"required"`
	Date                *string        `json:"date,omitempty"`
	Time                *string        `json:"time,omitempty"`
	Location            *string        `json:"location,omitempty"`
	IsFeatured          bool           `json:"is_featured"`
	PostType            types.PostType `json:"post_type" validate:"required"`
	UserID              string         `json:"user_id" validate:"required"`
	Images              []string       `json:"images,omitempty"`
}

type PostImageOutputDTO struct {
	ID       string `json:"id"`
	ImageURL string `json:"image_url"`
}

type CreatePostOutputDTO struct {
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Category            string                `json:"category"`
	SmallDescription    string                `json:"small_description"`
	CompleteDescription string                `json:"complete_description"`
	Date                *string               `json:"date,omitempty"`
	Time                *string               `json:"time,omitempty"`
	Location            *string               `json:"location,omitempty"`
	IsFeatured          bool                  `json:"is_featured"`
	PostType            types.PostType        `json:"post_type"`
	UserID              string                `json:"user_id"`
	Images              []*PostImageOutputDTO `json:"images,omitempty"`
	CreatedAt           time.Time             `json:"created_at"`
	UpdatedAt           time.Time             `json:"updated_at"`
}

type CreatePostUseCase struct {
	PostRepository post_repository.CreatePostRepository
	UserRepository user.GetUserByIdRepository
	StorageService storage.StorageService
}

func NewCreatePostUseCase(
	postRepository post_repository.CreatePostRepository,
	userRepository user.GetUserByIdRepository,
	storageService storage.StorageService,
) *CreatePostUseCase {
	return &CreatePostUseCase{
		PostRepository: postRepository,
		UserRepository: userRepository,
		StorageService: storageService,
	}
}

type imageUploadResult struct {
	index    int
	image    *entity.PostImage
	err      error
	imageURL string
}

func (uc *CreatePostUseCase) Execute(input CreatePostInputDTO) (*CreatePostOutputDTO, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userChan := make(chan struct {
		user *entity.User
		err  error
	}, 1)

	go func() {
		user, err := uc.UserRepository.GetUserById(input.UserID)
		userChan <- struct {
			user *entity.User
			err  error
		}{user, err}
	}()

	postID := uuid.New().String()
	post, err := entity.NewPost(
		postID,
		input.Title,
		input.Category,
		input.SmallDescription,
		input.CompleteDescription,
		input.Date,
		input.Time,
		input.Location,
		input.IsFeatured,
		input.PostType,
		input.UserID,
		nil,
	)
	if err != nil {
		cancel()
		return nil, err
	}

	userResult := <-userChan
	if userResult.err != nil {
		return nil, userResult.err
	}
	post.User = userResult.user

	var uploadedImages []*entity.PostImage
	var uploadedURLs []string

	if len(input.Images) > 0 {

		const maxWorkers = 5
		semaphore := make(chan struct{}, maxWorkers)
		resultChan := make(chan imageUploadResult, len(input.Images))
		var wg sync.WaitGroup

		for i, imageBase64 := range input.Images {
			wg.Add(1)
			go func(index int, imgBase64 string) {
				defer wg.Done()

				semaphore <- struct{}{}
				defer func() { <-semaphore }()

				imageData, err := base64.StdEncoding.DecodeString(imgBase64)
				if err != nil {
					resultChan <- imageUploadResult{
						index: index,
						err:   apperrors.NewValidationError("images", fmt.Sprintf("Imagem %d inválida em formato base64", index)),
					}
					return
				}

				filename := fmt.Sprintf("post_%s_img_%d_%s.jpg", postID, index, uuid.New().String())
				imageURL, err := uc.StorageService.UploadImageWithContext(ctx, imageData, filename, "posts")
				if err != nil {
					resultChan <- imageUploadResult{
						index: index,
						err:   apperrors.NewInternalError(fmt.Sprintf("Erro ao fazer upload da imagem %d", index), err),
					}
					return
				}

				resultChan <- imageUploadResult{
					index:    index,
					image:    entity.NewPostImage("", postID, imageURL),
					imageURL: imageURL,
					err:      nil,
				}
			}(i, imageBase64)
		}

		go func() {
			wg.Wait()
			close(resultChan)
		}()

		results := make([]imageUploadResult, len(input.Images))
		for result := range resultChan {
			results[result.index] = result
		}

		for _, result := range results {
			if result.err != nil {

				for _, url := range uploadedURLs {
					_ = uc.StorageService.DeleteImage(url, "posts")
				}
				return nil, result.err
			}
			uploadedImages = append(uploadedImages, result.image)
			uploadedURLs = append(uploadedURLs, result.imageURL)
		}

		post.Images = uploadedImages
	}

	postCreated, err := uc.PostRepository.CreatePost(post)
	if err != nil {

		for _, url := range uploadedURLs {
			_ = uc.StorageService.DeleteImage(url, "posts")
		}
		return nil, err
	}

	var imagesOutput []*PostImageOutputDTO
	if len(postCreated.Images) > 0 {
		imagesOutput = make([]*PostImageOutputDTO, len(postCreated.Images))
		for i, img := range postCreated.Images {
			imagesOutput[i] = &PostImageOutputDTO{
				ID:       img.ID,
				ImageURL: img.ImageURL,
			}
		}
	}

	return &CreatePostOutputDTO{
		ID:                  postCreated.ID,
		Title:               postCreated.Title,
		Category:            postCreated.Category,
		SmallDescription:    postCreated.SmallDescription,
		CompleteDescription: postCreated.CompleteDescription,
		Date:                postCreated.Date,
		Time:                postCreated.Time,
		Location:            postCreated.Location,
		IsFeatured:          postCreated.IsFeatured,
		PostType:            postCreated.PostType,
		UserID:              postCreated.UserID,
		Images:              imagesOutput,
		CreatedAt:           postCreated.CreatedAt,
		UpdatedAt:           postCreated.UpdatedAt,
	}, nil
}
