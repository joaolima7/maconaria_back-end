package timeline_usecase

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
)

type CreateTimelineInputDTO struct {
	Period      string `json:"period" validate:"required"`
	PdfData     string `json:"pdf_data" validate:"required"`
	IsHighlight bool   `json:"is_highlight"`
}

type CreateTimelineOutputDTO struct {
	ID          string    `json:"id"`
	Period      string    `json:"period"`
	PdfData     string    `json:"pdf_data"`
	IsHighlight bool      `json:"is_highlight"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTimelineUseCase struct {
	Repository timeline.CreateTimelineRepository
}

func NewCreateTimelineUseCase(repository timeline.CreateTimelineRepository) *CreateTimelineUseCase {
	return &CreateTimelineUseCase{
		Repository: repository,
	}
}

func (uc *CreateTimelineUseCase) Execute(input CreateTimelineInputDTO) (*CreateTimelineOutputDTO, error) {
	if input.PdfData == "" {
		return nil, apperrors.NewValidationError("PDF", "O PDF é obrigatório!")
	}

	pdfData, err := base64.StdEncoding.DecodeString(input.PdfData)
	if err != nil {
		return nil, apperrors.NewValidationError("PDF", "PDF em formato inválido!")
	}

	timelineEntity, err := entity.NewTimeline(
		uuid.New().String(),
		input.Period,
		pdfData,
		input.IsHighlight,
	)
	if err != nil {
		return nil, err
	}

	timelineCreated, err := uc.Repository.CreateTimeline(timelineEntity)
	if err != nil {
		return nil, err
	}

	pdfDataBase64 := ""
	if len(timelineCreated.PdfData) > 0 {
		pdfDataBase64 = base64.StdEncoding.EncodeToString(timelineCreated.PdfData)
	}

	return &CreateTimelineOutputDTO{
		ID:          timelineCreated.ID,
		Period:      timelineCreated.Period,
		PdfData:     pdfDataBase64,
		IsHighlight: timelineCreated.IsHighlight,
		CreatedAt:   timelineCreated.CreatedAt,
		UpdatedAt:   timelineCreated.UpdatedAt,
	}, nil
}
