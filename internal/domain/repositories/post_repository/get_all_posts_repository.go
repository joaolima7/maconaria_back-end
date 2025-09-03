package post_repository

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetAllPostsRepository interface {
	GetAllPosts() ([]*entity.Post, error)
}
