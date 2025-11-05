package timeline_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
)

type UpdateTimelineByIDInputDTO struct {
	ID          string `json:"id"`
	Period      string `json:"period" validate:"required"`
	PdfData     string `json:"pdf_data" validate:"required"`
	IsHighlight bool   `json:"is_highlight"`
}

type UpdateTimelineByIDOutputDTO struct {
	ID          string    `json:"id"`
	Period      string    `json:"period"`
	PdfData     string    `json:"pdf_data"`
	IsHighlight bool      `json:"is_highlight"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTimelineByIDUseCase struct {
	Repository timeline.UpdateTimelineByIDRepository
}

func NewUpdateTimelineByIDUseCase(repository timeline.UpdateTimelineByIDRepository) *UpdateTimelineByIDUseCase {
	return &UpdateTimelineByIDUseCase{
		Repository: repository,
	}
}

func (uc *UpdateTimelineByIDUseCase) Execute(input UpdateTimelineByIDInputDTO) (*UpdateTimelineByIDOutputDTO, error) {
	if input.PdfData == "" {
		return nil, apperrors.NewValidationError("PDF", "O PDF é obrigatório!")
	}

	pdfData, err := base64.StdEncoding.DecodeString(input.PdfData)
	if err != nil {
		return nil, apperrors.NewValidationError("PDF", "PDF em formato inválido!")
	}

	timelineEntity := &entity.Timeline{
		ID:          input.ID,
		Period:      input.Period,
		PdfData:     pdfData,
		IsHighlight: input.IsHighlight,
		UpdatedAt:   time.Now(),
	}

	if err := timelineEntity.Validate(); err != nil {
		return nil, err
	}

	timelineUpdated, err := uc.Repository.UpdateTimelineByID(timelineEntity)
	if err != nil {
		return nil, err
	}

	pdfDataBase64 := ""
	if len(timelineUpdated.PdfData) > 0 {
		pdfDataBase64 = base64.StdEncoding.EncodeToString(timelineUpdated.PdfData)
	}

	return &UpdateTimelineByIDOutputDTO{
		ID:          timelineUpdated.ID,
		Period:      timelineUpdated.Period,
		PdfData:     pdfDataBase64,
		IsHighlight: timelineUpdated.IsHighlight,
		CreatedAt:   timelineUpdated.CreatedAt,
		UpdatedAt:   timelineUpdated.UpdatedAt,
	}, nil
}
