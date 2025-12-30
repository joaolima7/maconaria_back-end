package wordkey_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/wordkey"
)

type UpdateWordKeyByIDInputDTO struct {
	ID      string `json:"id"`
	WordKey string `json:"wordkey" validate:"required,min=3"`
	Active  bool   `json:"active"`
}

type UpdateWordKeyByIDOutputDTO struct {
	ID        string    `json:"id"`
	WordKey   string    `json:"wordkey"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateWordKeyByIDUseCase struct {
	GetRepository        wordkey.GetWordKeyByIDRepository
	UpdateRepository     wordkey.UpdateWordKeyByIDRepository
	DeactivateRepository wordkey.DeactivateAllWordKeysRepository
}

func NewUpdateWordKeyByIDUseCase(
	getRepository wordkey.GetWordKeyByIDRepository,
	updateRepository wordkey.UpdateWordKeyByIDRepository,
	deactivateRepository wordkey.DeactivateAllWordKeysRepository,
) *UpdateWordKeyByIDUseCase {
	return &UpdateWordKeyByIDUseCase{
		GetRepository:        getRepository,
		UpdateRepository:     updateRepository,
		DeactivateRepository: deactivateRepository,
	}
}

func (uc *UpdateWordKeyByIDUseCase) Execute(input UpdateWordKeyByIDInputDTO) (*UpdateWordKeyByIDOutputDTO, error) {
	existingWordKey, err := uc.GetRepository.GetWordKeyByID(input.ID)
	if err != nil {
		return nil, err
	}

	// If setting this wordkey as active, deactivate all others first
	if input.Active && !existingWordKey.Active {
		if err := uc.DeactivateRepository.DeactivateAllWordKeys(); err != nil {
			return nil, err
		}
	}

	wordkeyEntity, err := entity.NewWordKey(
		input.ID,
		input.WordKey,
		input.Active,
	)
	if err != nil {
		return nil, err
	}

	wordkeyEntity.CreatedAt = existingWordKey.CreatedAt

	wordkeyUpdated, err := uc.UpdateRepository.UpdateWordKeyByID(wordkeyEntity)
	if err != nil {
		return nil, err
	}

	return &UpdateWordKeyByIDOutputDTO{
		ID:        wordkeyUpdated.ID,
		WordKey:   wordkeyUpdated.WordKey,
		Active:    wordkeyUpdated.Active,
		CreatedAt: wordkeyUpdated.CreatedAt,
	}, nil
}
