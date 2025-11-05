package acacia

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetAllAcaciasRepository interface {
	GetAllAcacias() ([]*entity.Acacia, error)
}
