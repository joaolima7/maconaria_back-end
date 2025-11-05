package worker_usecase

import "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/worker"

type DeleteWorkerInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeleteWorkerUseCase struct {
	Repository worker.DeleteWorkerRepository
}

func NewDeleteWorkerUseCase(repository worker.DeleteWorkerRepository) *DeleteWorkerUseCase {
	return &DeleteWorkerUseCase{
		Repository: repository,
	}
}

func (uc *DeleteWorkerUseCase) Execute(input DeleteWorkerInputDTO) error {
	return uc.Repository.DeleteWorker(input.ID)
}
