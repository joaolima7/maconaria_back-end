package timeline_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
)

type GetTimelineByIDInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type GetTimelineByIDOutputDTO struct {
	ID          string    `json:"id"`
	Period      string    `json:"period"`
	PdfURL      string    `json:"pdf_url"`
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

	return &GetTimelineByIDOutputDTO{
		ID:          t.ID,
		Period:      t.Period,
		PdfURL:      t.PdfURL,
		IsHighlight: t.IsHighlight,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}, nil
}
