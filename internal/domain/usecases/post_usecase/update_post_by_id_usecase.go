package post_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	post_repository "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/post"
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
	Repository post_repository.UpdatePostByIDRepository
}

func NewUpdatePostByIDUseCase(repository post_repository.UpdatePostByIDRepository) *UpdatePostByIDUseCase {
	return &UpdatePostByIDUseCase{
		Repository: repository,
	}
}

func (uc *UpdatePostByIDUseCase) Execute(input UpdatePostByIDInputDTO) (*UpdatePostByIDOutputDTO, error) {
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

			post.Images[i] = entity.NewPostImage("", post.ID, imageData)
		}
	}

	updatedPost, err := uc.Repository.UpdatePostByID(post)
	if err != nil {
		return nil, err
	}

	var imagesOutput []*PostImageOutputDTO
	if len(updatedPost.Images) > 0 {
		imagesOutput = make([]*PostImageOutputDTO, len(updatedPost.Images))
		for i, img := range updatedPost.Images {
			imagesOutput[i] = &PostImageOutputDTO{
				ID:        img.ID,
				ImageData: base64.StdEncoding.EncodeToString(img.ImageData),
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
