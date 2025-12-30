package entity

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
)

type WordKey struct {
	ID        string
	WordKey   string
	Active    bool
	CreatedAt time.Time
}

func NewWordKey(
	id string,
	wordkey string,
	active bool,
) (*WordKey, error) {
	wordkeyEntity := &WordKey{
		ID:        id,
		WordKey:   wordkey,
		Active:    active,
		CreatedAt: time.Now(),
	}

	if err := wordkeyEntity.Validate(); err != nil {
		return nil, err
	}

	return wordkeyEntity, nil
}

func (w *WordKey) Validate() error {
	if err := w.ValidateWordKey(); err != nil {
		return err
	}
	return nil
}

func (w *WordKey) ValidateWordKey() error {
	if len(w.WordKey) == 0 {
		return apperrors.NewValidationError("wordkey", "A palavra chave não pode ser vazia!")
	}
	if len(w.WordKey) < 3 {
		return apperrors.NewValidationError("wordkey", "A palavra chave deve ter no mínimo 3 caracteres!")
	}
	return nil
}

func (w *WordKey) Activate() {
	w.Active = true
}

func (w *WordKey) Deactivate() {
	w.Active = false
}
