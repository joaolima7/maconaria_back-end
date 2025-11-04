package user

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetUserByEmailRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
}
