package library_usecase

import "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"

type DeleteLibraryInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeleteLibraryUseCase struct {
	Repository library.DeleteLibraryRepository
}

func NewDeleteLibraryUseCase(repository library.DeleteLibraryRepository) *DeleteLibraryUseCase {
	return &DeleteLibraryUseCase{
		Repository: repository,
	}
}

func (uc *DeleteLibraryUseCase) Execute(input DeleteLibraryInputDTO) error {
	return uc.Repository.DeleteLibrary(input.ID)
}
