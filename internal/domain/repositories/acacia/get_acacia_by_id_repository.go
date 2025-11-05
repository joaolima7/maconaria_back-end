package acacia

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetAcaciaByIDRepository interface {
	GetAcaciaByID(id string) (*entity.Acacia, error)
}
