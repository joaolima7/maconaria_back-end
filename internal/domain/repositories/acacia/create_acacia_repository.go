package acacia

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type CreateAcaciaRepository interface {
	CreateAcacia(acacia *entity.Acacia) (*entity.Acacia, error)
}
