package user_repository

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetAllUsersRepository interface {
	GetAllUsers() ([]*entity.User, error)
}
