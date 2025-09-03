package user_repository

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type UpdateUserByIDRepository interface {
	UpdateUserByID(user *entity.User) (*entity.User, error)
}
