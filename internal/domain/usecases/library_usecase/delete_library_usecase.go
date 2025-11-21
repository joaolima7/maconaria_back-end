package library_usecase

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type DeleteLibraryInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeleteLibraryUseCase struct {
	Repository     library.DeleteLibraryRepository
	GetRepository  library.GetLibraryByIDRepository
	StorageService storage.StorageService
}

func NewDeleteLibraryUseCase(
	repository library.DeleteLibraryRepository,
	getRepository library.GetLibraryByIDRepository,
	storageService storage.StorageService,
) *DeleteLibraryUseCase {
	return &DeleteLibraryUseCase{
		Repository:     repository,
		GetRepository:  getRepository,
		StorageService: storageService,
	}
}

func (uc *DeleteLibraryUseCase) Execute(input DeleteLibraryInputDTO) error {

	l, err := uc.GetRepository.GetLibraryByID(input.ID)
	if err != nil {
		return err
	}

	if err := uc.Repository.DeleteLibrary(input.ID); err != nil {
		return err
	}

	if l.FileURL != "" {
		_ = uc.StorageService.DeletePDF(l.FileURL, "libraries")
	}

	if l.CoverURL != "" {
		_ = uc.StorageService.DeleteImage(l.CoverURL, "libraries")
	}

	return nil
}
