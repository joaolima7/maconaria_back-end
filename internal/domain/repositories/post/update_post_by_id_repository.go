package post_repository

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type UpdatePostByIDRepository interface {
	UpdatePostByID(post *entity.Post) (*entity.Post, error)
}
