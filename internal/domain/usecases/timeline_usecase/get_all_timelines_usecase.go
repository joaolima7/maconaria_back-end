package timeline_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
)

type GetAllTimelinesOutputDTO struct {
	ID          string    `json:"id"`
	Period      string    `json:"period"`
	PdfData     string    `json:"pdf_data"`
	IsHighlight bool      `json:"is_highlight"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllTimelinesUseCase struct {
	Repository timeline.GetAllTimelinesRepository
}

func NewGetAllTimelinesUseCase(repository timeline.GetAllTimelinesRepository) *GetAllTimelinesUseCase {
	return &GetAllTimelinesUseCase{
		Repository: repository,
	}
}

func (uc *GetAllTimelinesUseCase) Execute() ([]*GetAllTimelinesOutputDTO, error) {
	timelines, err := uc.Repository.GetAllTimelines()
	if err != nil {
		return nil, err
	}

	output := make([]*GetAllTimelinesOutputDTO, len(timelines))
	for i, t := range timelines {
		pdfData := ""
		if len(t.PdfData) > 0 {
			pdfData = base64.StdEncoding.EncodeToString(t.PdfData)
		}

		output[i] = &GetAllTimelinesOutputDTO{
			ID:          t.ID,
			Period:      t.Period,
			PdfData:     pdfData,
			IsHighlight: t.IsHighlight,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		}
	}

	return output, nil
}
