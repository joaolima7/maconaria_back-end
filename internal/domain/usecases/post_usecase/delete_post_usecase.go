package post_usecase

import (
	post_repository "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/post"
)

type DeletePostInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeletePostUseCase struct {
	Repository post_repository.DeletePostRepository
}

func NewDeletePostUseCase(repository post_repository.DeletePostRepository) *DeletePostUseCase {
	return &DeletePostUseCase{
		Repository: repository,
	}
}

func (uc *DeletePostUseCase) Execute(input DeletePostInputDTO) error {
	return uc.Repository.Delete(input.ID)
}
