package entity

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
)

type Timeline struct {
	ID          string
	Period      string
	PdfData     []byte
	IsHighlight bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTimeline(
	id string,
	period string,
	pdfData []byte,
	isHighlight bool,
) (*Timeline, error) {
	timeline := &Timeline{
		ID:          id,
		Period:      period,
		PdfData:     pdfData,
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
	if err := t.ValidatePdf(); err != nil {
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

func (t *Timeline) ValidatePdf() error {
	if len(t.PdfData) == 0 {
		return apperrors.NewValidationError("PDF", "O PDF é obrigatório!")
	}
	return nil
}

func (t *Timeline) UpdatePdf(pdfData []byte) {
	t.PdfData = pdfData
	t.UpdatedAt = time.Now()
}

func (t *Timeline) ToggleHighlight() {
	t.IsHighlight = !t.IsHighlight
	t.UpdatedAt = time.Now()
}
