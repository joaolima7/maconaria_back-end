package entity

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
)

type Timeline struct {
	ID          string
	Period      string
	PdfURL      string
	IsHighlight bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTimeline(
	id string,
	period string,
	pdfURL string,
	isHighlight bool,
) (*Timeline, error) {
	timeline := &Timeline{
		ID:          id,
		Period:      period,
		PdfURL:      pdfURL,
		IsHighlight: isHighlight,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := timeline.Validate(); err != nil {
		return nil, err
	}

	return timeline, nil
}

func (t *Timeline) Validate() error {
	if err := t.ValidatePeriod(); err != nil {
		return err
	}
	if err := t.ValidatePdfURL(); err != nil {
		return err
	}
	return nil
}

func (t *Timeline) ValidatePeriod() error {
	if len(t.Period) == 0 {
		return apperrors.NewValidationError("período", "O período não pode ser vazio!")
	}
	if len(t.Period) > 50 {
		return apperrors.NewValidationError("período", "O período deve ter no máximo 50 caracteres!")
	}
	return nil
}

func (t *Timeline) ValidatePdfURL() error {
	if t.PdfURL == "" {
		return apperrors.NewValidationError("PDF", "A URL do PDF é obrigatória!")
	}
	return nil
}

func (t *Timeline) UpdatePdf(pdfURL string) {
	t.PdfURL = pdfURL
	t.UpdatedAt = time.Now()
}

func (t *Timeline) ToggleHighlight() {
	t.IsHighlight = !t.IsHighlight
	t.UpdatedAt = time.Now()
}
