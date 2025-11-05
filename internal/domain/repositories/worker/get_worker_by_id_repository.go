package worker

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetWorkerByIDRepository interface {
	GetWorkerByID(id string) (*entity.Worker, error)
}
