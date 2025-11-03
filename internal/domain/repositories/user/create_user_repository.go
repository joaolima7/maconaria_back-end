package user

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type CreateUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
}
