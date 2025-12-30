package wordkey_usecase

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/wordkey"
)

type DeleteWordKeyInputDTO struct {
	ID string `json:"id"`
}

type DeleteWordKeyUseCase struct {
	GetRepository    wordkey.GetWordKeyByIDRepository
	DeleteRepository wordkey.DeleteWordKeyRepository
}

func NewDeleteWordKeyUseCase(
	getRepository wordkey.GetWordKeyByIDRepository,
	deleteRepository wordkey.DeleteWordKeyRepository,
) *DeleteWordKeyUseCase {
	return &DeleteWordKeyUseCase{
		GetRepository:    getRepository,
		DeleteRepository: deleteRepository,
	}
}

func (uc *DeleteWordKeyUseCase) Execute(input DeleteWordKeyInputDTO) error {
	_, err := uc.GetRepository.GetWordKeyByID(input.ID)
	if err != nil {
		return err
	}

	if err := uc.DeleteRepository.DeleteWordKey(input.ID); err != nil {
		return err
	}

	return nil
}
