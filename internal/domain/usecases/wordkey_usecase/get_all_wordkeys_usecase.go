package wordkey_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/wordkey"
)

type GetAllWordKeysOutputDTO struct {
	ID        string    `json:"id"`
	WordKey   string    `json:"wordkey"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllWordKeysUseCase struct {
	Repository wordkey.GetAllWordKeysRepository
}

func NewGetAllWordKeysUseCase(repository wordkey.GetAllWordKeysRepository) *GetAllWordKeysUseCase {
	return &GetAllWordKeysUseCase{
		Repository: repository,
	}
}

func (uc *GetAllWordKeysUseCase) Execute() ([]*GetAllWordKeysOutputDTO, error) {
	wordkeys, err := uc.Repository.GetAllWordKeys()
	if err != nil {
		return nil, err
	}

	output := make([]*GetAllWordKeysOutputDTO, len(wordkeys))
	for i, wk := range wordkeys {
		output[i] = &GetAllWordKeysOutputDTO{
			ID:        wk.ID,
			WordKey:   wk.WordKey,
			Active:    wk.Active,
			CreatedAt: wk.CreatedAt,
		}
	}

	return output, nil
}
