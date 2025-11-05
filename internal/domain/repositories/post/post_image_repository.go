package post_repository

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type PostImageRepository interface {
	CreatePostImage(image *entity.PostImage) error
	GetPostImages(postID string) ([]*entity.PostImage, error)
	DeletePostImagesByPostID(postID string) error
	DeletePostImageByID(imageID string) error
}
