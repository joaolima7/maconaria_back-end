package worker

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetAllWorkersRepository interface {
	GetAllWorkers() ([]*entity.Worker, error)
}
