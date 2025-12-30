package wordkey_usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/wordkey"
)

type CreateWordKeyInputDTO struct {
	WordKey string `json:"wordkey" validate:"required"`
	Active  bool   `json:"active"`
}

type CreateWordKeyOutputDTO struct {
	ID        string    `json:"id"`
	WordKey   string    `json:"wordkey"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateWordKeyUseCase struct {
	CreateRepository     wordkey.CreateWordKeyRepository
	DeactivateRepository wordkey.DeactivateAllWordKeysRepository
}

func NewCreateWordKeyUseCase(
	createRepository wordkey.CreateWordKeyRepository,
	deactivateRepository wordkey.DeactivateAllWordKeysRepository,
) *CreateWordKeyUseCase {
	return &CreateWordKeyUseCase{
		CreateRepository:     createRepository,
		DeactivateRepository: deactivateRepository,
	}
}

func (uc *CreateWordKeyUseCase) Execute(input CreateWordKeyInputDTO) (*CreateWordKeyOutputDTO, error) {
	// If creating an active wordkey, deactivate all others first
	if input.Active {
		if err := uc.DeactivateRepository.DeactivateAllWordKeys(); err != nil {
			return nil, err
		}
	}

	wordkeyID := uuid.New().String()

	wordkeyEntity, err := entity.NewWordKey(
		wordkeyID,
		input.WordKey,
		input.Active,
	)
	if err != nil {
		return nil, err
	}

	wordkeyCreated, err := uc.CreateRepository.CreateWordKey(wordkeyEntity)
	if err != nil {
		return nil, apperrors.NewInternalError("Erro ao criar palavra chave!", err)
	}

	return &CreateWordKeyOutputDTO{
		ID:        wordkeyCreated.ID,
		WordKey:   wordkeyCreated.WordKey,
		Active:    wordkeyCreated.Active,
		CreatedAt: wordkeyCreated.CreatedAt,
	}, nil
}
