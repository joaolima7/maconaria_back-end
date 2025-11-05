package worker

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type CreateWorkerRepository interface {
	CreateWorker(worker *entity.Worker) (*entity.Worker, error)
}
