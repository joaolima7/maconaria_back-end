package timeline_usecase

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type CreateTimelineInputDTO struct {
	Period      string `json:"period" validate:"required"`
	PdfData     string `json:"pdf_data" validate:"required,base64"`
	IsHighlight bool   `json:"is_highlight"`
}

type CreateTimelineOutputDTO struct {
	ID          string    `json:"id"`
	Period      string    `json:"period"`
	PdfURL      string    `json:"pdf_url"`
	IsHighlight bool      `json:"is_highlight"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTimelineUseCase struct {
	Repository     timeline.CreateTimelineRepository
	StorageService storage.StorageService
}

func NewCreateTimelineUseCase(
	repository timeline.CreateTimelineRepository,
	storageService storage.StorageService,
) *CreateTimelineUseCase {
	return &CreateTimelineUseCase{
		Repository:     repository,
		StorageService: storageService,
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

	timelineID := uuid.New().String()
	filename := fmt.Sprintf("timeline_%s_%s.pdf", timelineID, uuid.New().String())

	pdfURL, err := uc.StorageService.UploadPDF(pdfData, filename, "timelines")
	if err != nil {
		return nil, apperrors.NewInternalError("Erro ao fazer upload do PDF", err)
	}

	timelineEntity, err := entity.NewTimeline(
		timelineID,
		input.Period,
		pdfURL,
		input.IsHighlight,
	)
	if err != nil {

		_ = uc.StorageService.DeletePDF(pdfURL, "timelines")
		return nil, err
	}

	timelineCreated, err := uc.Repository.CreateTimeline(timelineEntity)
	if err != nil {

		_ = uc.StorageService.DeletePDF(pdfURL, "timelines")
		return nil, err
	}

	return &CreateTimelineOutputDTO{
		ID:          timelineCreated.ID,
		Period:      timelineCreated.Period,
		PdfURL:      timelineCreated.PdfURL,
		IsHighlight: timelineCreated.IsHighlight,
		CreatedAt:   timelineCreated.CreatedAt,
		UpdatedAt:   timelineCreated.UpdatedAt,
	}, nil
}
