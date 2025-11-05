package acacia

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type UpdateAcaciaByIDRepository interface {
	UpdateAcaciaByID(acacia *entity.Acacia) (*entity.Acacia, error)
}
