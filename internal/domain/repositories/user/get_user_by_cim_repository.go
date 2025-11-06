package user

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetUserByCIMRepository interface {
	GetUserByCIM(cim string) (*entity.User, error)
}
