package wordkey_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/wordkey"
)

type GetWordKeyByIDInputDTO struct {
	ID string `json:"id"`
}

type GetWordKeyByIDOutputDTO struct {
	ID        string    `json:"id"`
	WordKey   string    `json:"wordkey"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type GetWordKeyByIDUseCase struct {
	Repository wordkey.GetWordKeyByIDRepository
}

func NewGetWordKeyByIDUseCase(repository wordkey.GetWordKeyByIDRepository) *GetWordKeyByIDUseCase {
	return &GetWordKeyByIDUseCase{
		Repository: repository,
	}
}

func (uc *GetWordKeyByIDUseCase) Execute(input GetWordKeyByIDInputDTO) (*GetWordKeyByIDOutputDTO, error) {
	wordkeyFound, err := uc.Repository.GetWordKeyByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetWordKeyByIDOutputDTO{
		ID:        wordkeyFound.ID,
		WordKey:   wordkeyFound.WordKey,
		Active:    wordkeyFound.Active,
		CreatedAt: wordkeyFound.CreatedAt,
	}, nil
}
