package acacia_usecase

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/acacia"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type DeleteAcaciaInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeleteAcaciaUseCase struct {
	Repository     acacia.DeleteAcaciaRepository
	GetRepository  acacia.GetAcaciaByIDRepository
	StorageService storage.StorageService
}

func NewDeleteAcaciaUseCase(
	repository acacia.DeleteAcaciaRepository,
	getRepository acacia.GetAcaciaByIDRepository,
	storageService storage.StorageService,
) *DeleteAcaciaUseCase {
	return &DeleteAcaciaUseCase{
		Repository:     repository,
		GetRepository:  getRepository,
		StorageService: storageService,
	}
}

func (uc *DeleteAcaciaUseCase) Execute(input DeleteAcaciaInputDTO) error {

	a, err := uc.GetRepository.GetAcaciaByID(input.ID)
	if err != nil {
		return err
	}

	if err := uc.Repository.DeleteAcacia(input.ID); err != nil {
		return err
	}

	_ = uc.StorageService.DeleteImage(a.ImageURL, "acacias")

	return nil
}
