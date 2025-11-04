package user

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetUserByIdRepository interface {
	GetUserById(id string) (*entity.User, error)
}
