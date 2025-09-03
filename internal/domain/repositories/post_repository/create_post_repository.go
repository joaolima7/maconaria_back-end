package post_repository

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type CreatePostRepository interface {
	CreatePost(post *entity.Post) (*entity.Post, error)
}
