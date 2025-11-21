package timeline_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
)

type GetAllTimelinesOutputDTO struct {
	ID          string    `json:"id"`
	Period      string    `json:"period"`
	PdfURL      string    `json:"pdf_url"`
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
		output[i] = &GetAllTimelinesOutputDTO{
			ID:          t.ID,
			Period:      t.Period,
			PdfURL:      t.PdfURL,
			IsHighlight: t.IsHighlight,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		}
	}

	return output, nil
}
