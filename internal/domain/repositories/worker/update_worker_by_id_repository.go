package worker

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type UpdateWorkerByIDRepository interface {
	UpdateWorkerByID(worker *entity.Worker) (*entity.Worker, error)
}
