package post_usecase

import (
	"encoding/base64"
	"time"

	post_repository "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/post"
	"github.com/joaolima7/maconaria_back-end/internal/types"
)

type GetAllPostsOutputDTO struct {
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

type GetAllPostsUseCase struct {
	Repository post_repository.GetAllPostsRepository
}

func NewGetAllPostsUseCase(repository post_repository.GetAllPostsRepository) *GetAllPostsUseCase {
	return &GetAllPostsUseCase{
		Repository: repository,
	}
}

func (uc *GetAllPostsUseCase) Execute() ([]*GetAllPostsOutputDTO, error) {
	posts, err := uc.Repository.GetAllPosts()
	if err != nil {
		return nil, err
	}

	output := make([]*GetAllPostsOutputDTO, len(posts))
	for i, post := range posts {

		var imagesOutput []*PostImageOutputDTO
		if len(post.Images) > 0 {
			imagesOutput = make([]*PostImageOutputDTO, len(post.Images))
			for j, img := range post.Images {
				imagesOutput[j] = &PostImageOutputDTO{
					ID:        img.ID,
					ImageData: base64.StdEncoding.EncodeToString(img.ImageData),
				}
			}
		}

		output[i] = &GetAllPostsOutputDTO{
			ID:                  post.ID,
			Title:               post.Title,
			Category:            post.Category,
			SmallDescription:    post.SmallDescription,
			CompleteDescription: post.CompleteDescription,
			Date:                post.Date,
			Time:                post.Time,
			Location:            post.Location,
			IsFeatured:          post.IsFeatured,
			PostType:            post.PostType,
			UserID:              post.UserID,
			Images:              imagesOutput,
			CreatedAt:           post.CreatedAt,
			UpdatedAt:           post.UpdatedAt,
		}
	}

	return output, nil
}
