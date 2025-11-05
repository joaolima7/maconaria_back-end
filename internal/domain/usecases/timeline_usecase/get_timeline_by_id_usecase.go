package timeline_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
)

type GetTimelineByIDInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type GetTimelineByIDOutputDTO struct {
	ID          string    `json:"id"`
	Period      string    `json:"period"`
	PdfData     string    `json:"pdf_data"`
	IsHighlight bool      `json:"is_highlight"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetTimelineByIDUseCase struct {
	Repository timeline.GetTimelineByIDRepository
}

func NewGetTimelineByIDUseCase(repository timeline.GetTimelineByIDRepository) *GetTimelineByIDUseCase {
	return &GetTimelineByIDUseCase{
		Repository: repository,
	}
}

func (uc *GetTimelineByIDUseCase) Execute(input GetTimelineByIDInputDTO) (*GetTimelineByIDOutputDTO, error) {
	t, err := uc.Repository.GetTimelineByID(input.ID)
	if err != nil {
		return nil, err
	}

	pdfData := ""
	if len(t.PdfData) > 0 {
		pdfData = base64.StdEncoding.EncodeToString(t.PdfData)
	}

	return &GetTimelineByIDOutputDTO{
		ID:          t.ID,
		Period:      t.Period,
		PdfData:     pdfData,
		IsHighlight: t.IsHighlight,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}, nil
}
