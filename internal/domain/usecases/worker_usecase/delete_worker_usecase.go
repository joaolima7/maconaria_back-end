package worker_usecase

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/worker"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type DeleteWorkerInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeleteWorkerUseCase struct {
	Repository     worker.DeleteWorkerRepository
	GetRepository  worker.GetWorkerByIDRepository
	StorageService storage.StorageService
}

func NewDeleteWorkerUseCase(
	repository worker.DeleteWorkerRepository,
	getRepository worker.GetWorkerByIDRepository,
	storageService storage.StorageService,
) *DeleteWorkerUseCase {
	return &DeleteWorkerUseCase{
		Repository:     repository,
		GetRepository:  getRepository,
		StorageService: storageService,
	}
}

func (uc *DeleteWorkerUseCase) Execute(input DeleteWorkerInputDTO) error {

	w, err := uc.GetRepository.GetWorkerByID(input.ID)
	if err != nil {
		return err
	}

	if err := uc.Repository.DeleteWorker(input.ID); err != nil {
		return err
	}

	_ = uc.StorageService.DeleteImage(w.ImageURL, "workers")

	return nil
}
