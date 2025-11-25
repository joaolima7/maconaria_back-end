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
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
	"github.com/joaolima7/maconaria_back-end/internal/types"
)

type UpdatePostByIDInputDTO struct {
	ID                  string         `json:"id"`
	Title               string         `json:"title" validate:"required"`
	Category            string         `json:"category" validate:"required"`
	SmallDescription    string         `json:"small_description" validate:"required"`
	CompleteDescription string         `json:"complete_description" validate:"required"`
	Date                *string        `json:"date,omitempty"`
	Time                *string        `json:"time,omitempty"`
	Location            *string        `json:"location,omitempty"`
	IsFeatured          bool           `json:"is_featured"`
	PostType            types.PostType `json:"post_type" validate:"required"`
	Images              *[]string      `json:"images,omitempty"`
}

type UpdatePostByIDOutputDTO struct {
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

type UpdatePostByIDUseCase struct {
	Repository      post_repository.UpdatePostByIDRepository
	ImageRepository post_repository.PostImageRepository
	StorageService  storage.StorageService
}

func NewUpdatePostByIDUseCase(
	repository post_repository.UpdatePostByIDRepository,
	imageRepository post_repository.PostImageRepository,
	storageService storage.StorageService,
) *UpdatePostByIDUseCase {
	return &UpdatePostByIDUseCase{
		Repository:      repository,
		ImageRepository: imageRepository,
		StorageService:  storageService,
	}
}

func (uc *UpdatePostByIDUseCase) Execute(input UpdatePostByIDInputDTO) (*UpdatePostByIDOutputDTO, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	oldImages, err := uc.ImageRepository.GetPostImages(input.ID)
	if err != nil {
		return nil, err
	}

	post := &entity.Post{
		ID:                  input.ID,
		Title:               input.Title,
		Category:            input.Category,
		SmallDescription:    input.SmallDescription,
		CompleteDescription: input.CompleteDescription,
		Date:                input.Date,
		Time:                input.Time,
		Location:            input.Location,
		IsFeatured:          input.IsFeatured,
		PostType:            input.PostType,
		UpdatedAt:           time.Now(),
	}

	var shouldUpdateImages bool
	var newImages []*entity.PostImage
	var uploadedURLs []string

	if input.Images == nil {

		shouldUpdateImages = false
		post.Images = oldImages
	} else if len(*input.Images) == 0 {

		shouldUpdateImages = true
		post.Images = []*entity.PostImage{}
	} else {

		shouldUpdateImages = true

		const maxWorkers = 5
		semaphore := make(chan struct{}, maxWorkers)
		resultChan := make(chan imageUploadResult, len(*input.Images))
		var wg sync.WaitGroup

		for i, imageBase64 := range *input.Images {
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

				filename := fmt.Sprintf("post_%s_img_%d_%s.jpg", post.ID, index, uuid.New().String())
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
					image:    entity.NewPostImage("", post.ID, imageURL),
					imageURL: imageURL,
					err:      nil,
				}
			}(i, imageBase64)
		}

		go func() {
			wg.Wait()
			close(resultChan)
		}()

		results := make([]imageUploadResult, len(*input.Images))
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
			newImages = append(newImages, result.image)
			uploadedURLs = append(uploadedURLs, result.imageURL)
		}

		post.Images = newImages
	}

	updatedPost, err := uc.Repository.UpdatePostByID(post)
	if err != nil {

		if shouldUpdateImages {
			for _, url := range uploadedURLs {
				_ = uc.StorageService.DeleteImage(url, "posts")
			}
		}
		return nil, err
	}

	if shouldUpdateImages && len(oldImages) > 0 {
		for _, oldImg := range oldImages {
			_ = uc.StorageService.DeleteImage(oldImg.ImageURL, "posts")
		}
	}

	var imagesOutput []*PostImageOutputDTO
	if len(updatedPost.Images) > 0 {
		imagesOutput = make([]*PostImageOutputDTO, len(updatedPost.Images))
		for i, img := range updatedPost.Images {
			imagesOutput[i] = &PostImageOutputDTO{
				ID:       img.ID,
				ImageURL: img.ImageURL,
			}
		}
	}

	return &UpdatePostByIDOutputDTO{
		ID:                  updatedPost.ID,
		Title:               updatedPost.Title,
		Category:            updatedPost.Category,
		SmallDescription:    updatedPost.SmallDescription,
		CompleteDescription: updatedPost.CompleteDescription,
		Date:                updatedPost.Date,
		Time:                updatedPost.Time,
		Location:            updatedPost.Location,
		IsFeatured:          updatedPost.IsFeatured,
		PostType:            updatedPost.PostType,
		UserID:              updatedPost.UserID,
		Images:              imagesOutput,
		CreatedAt:           updatedPost.CreatedAt,
		UpdatedAt:           updatedPost.UpdatedAt,
	}, nil
}
