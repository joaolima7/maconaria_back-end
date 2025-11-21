package post_usecase

import (
	"encoding/base64"
	"fmt"
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
	Images              []string       `json:"images,omitempty"`
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

	if len(input.Images) > 0 {
		post.Images = make([]*entity.PostImage, len(input.Images))
		for i, imageBase64 := range input.Images {
			imageData, err := base64.StdEncoding.DecodeString(imageBase64)
			if err != nil {
				return nil, apperrors.NewValidationError("images", "Imagem inválida em formato base64")
			}

			filename := fmt.Sprintf("post_%s_img_%d_%s.jpg", post.ID, i, uuid.New().String())

			imageURL, err := uc.StorageService.UploadImage(imageData, filename, "posts")
			if err != nil {
				return nil, apperrors.NewInternalError("Erro ao fazer upload da imagem", err)
			}

			post.Images[i] = entity.NewPostImage("", post.ID, imageURL)
		}
	}

	updatedPost, err := uc.Repository.UpdatePostByID(post)
	if err != nil {

		for _, img := range post.Images {
			_ = uc.StorageService.DeleteImage(img.ImageURL, "posts")
		}
		return nil, err
	}

	for _, oldImg := range oldImages {
		_ = uc.StorageService.DeleteImage(oldImg.ImageURL, "posts")
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
