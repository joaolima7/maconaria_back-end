package acacia_usecase

import "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/acacia"

type DeleteAcaciaInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeleteAcaciaUseCase struct {
	Repository acacia.DeleteAcaciaRepository
}

func NewDeleteAcaciaUseCase(repository acacia.DeleteAcaciaRepository) *DeleteAcaciaUseCase {
	return &DeleteAcaciaUseCase{
		Repository: repository,
	}
}

func (uc *DeleteAcaciaUseCase) Execute(input DeleteAcaciaInputDTO) error {
	return uc.Repository.DeleteAcacia(input.ID)
}
