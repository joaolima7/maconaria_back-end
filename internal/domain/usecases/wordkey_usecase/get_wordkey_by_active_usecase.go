package wordkey_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/wordkey"
)

type GetWordKeyByActiveOutputDTO struct {
	ID        string    `json:"id"`
	WordKey   string    `json:"wordkey"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type GetWordKeyByActiveUseCase struct {
	Repository wordkey.GetWordKeyByActiveRepository
}

func NewGetWordKeyByActiveUseCase(repository wordkey.GetWordKeyByActiveRepository) *GetWordKeyByActiveUseCase {
	return &GetWordKeyByActiveUseCase{
		Repository: repository,
	}
}

func (uc *GetWordKeyByActiveUseCase) Execute() (*GetWordKeyByActiveOutputDTO, error) {
	wordkeyFound, err := uc.Repository.GetWordKeyByActive()
	if err != nil {
		return nil, err
	}

	return &GetWordKeyByActiveOutputDTO{
		ID:        wordkeyFound.ID,
		WordKey:   wordkeyFound.WordKey,
		Active:    wordkeyFound.Active,
		CreatedAt: wordkeyFound.CreatedAt,
	}, nil
}
