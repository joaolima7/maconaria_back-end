package post_usecase

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	post_repository "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/post"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
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
	ID        string `json:"id"`
	ImageData string `json:"image_data"`
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
}

func NewCreatePostUseCase(
	postRepository post_repository.CreatePostRepository,
	userRepository user.GetUserByIdRepository,
) *CreatePostUseCase {
	return &CreatePostUseCase{
		PostRepository: postRepository,
		UserRepository: userRepository,
	}
}

func (uc *CreatePostUseCase) Execute(input CreatePostInputDTO) (*CreatePostOutputDTO, error) {
	user, err := uc.UserRepository.GetUserById(input.UserID)
	if err != nil {
		return nil, err
	}

	post, err := entity.NewPost(
		uuid.New().String(),
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
		user,
	)
	if err != nil {
		return nil, err
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

	postCreated, err := uc.PostRepository.CreatePost(post)
	if err != nil {
		return nil, err
	}

	var imagesOutput []*PostImageOutputDTO
	if len(postCreated.Images) > 0 {
		imagesOutput = make([]*PostImageOutputDTO, len(postCreated.Images))
		for i, img := range postCreated.Images {
			imagesOutput[i] = &PostImageOutputDTO{
				ID:        img.ID,
				ImageData: base64.StdEncoding.EncodeToString(img.ImageData),
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
