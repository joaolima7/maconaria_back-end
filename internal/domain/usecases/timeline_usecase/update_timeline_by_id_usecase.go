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

type UpdateTimelineByIDInputDTO struct {
	ID          string `json:"id"`
	Period      string `json:"period" validate:"required"`
	PdfData     string `json:"pdf_data,omitempty"`
	IsHighlight bool   `json:"is_highlight"`
}

type UpdateTimelineByIDOutputDTO struct {
	ID          string    `json:"id"`
	Period      string    `json:"period"`
	PdfURL      string    `json:"pdf_url"`
	IsHighlight bool      `json:"is_highlight"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTimelineByIDUseCase struct {
	Repository     timeline.UpdateTimelineByIDRepository
	GetRepository  timeline.GetTimelineByIDRepository
	StorageService storage.StorageService
}

func NewUpdateTimelineByIDUseCase(
	repository timeline.UpdateTimelineByIDRepository,
	getRepository timeline.GetTimelineByIDRepository,
	storageService storage.StorageService,
) *UpdateTimelineByIDUseCase {
	return &UpdateTimelineByIDUseCase{
		Repository:     repository,
		GetRepository:  getRepository,
		StorageService: storageService,
	}
}

func (uc *UpdateTimelineByIDUseCase) Execute(input UpdateTimelineByIDInputDTO) (*UpdateTimelineByIDOutputDTO, error) {

	existingTimeline, err := uc.GetRepository.GetTimelineByID(input.ID)
	if err != nil {
		return nil, err
	}

	pdfURL := existingTimeline.PdfURL
	var oldPdfURL string

	if input.PdfData != "" {
		pdfData, err := base64.StdEncoding.DecodeString(input.PdfData)
		if err != nil {
			return nil, apperrors.NewValidationError("PDF", "PDF em formato inválido!")
		}

		filename := fmt.Sprintf("timeline_%s_%s.pdf", input.ID, uuid.New().String())

		newPdfURL, err := uc.StorageService.UploadPDF(pdfData, filename, "timelines")
		if err != nil {
			return nil, apperrors.NewInternalError("Erro ao fazer upload do PDF", err)
		}

		oldPdfURL = pdfURL
		pdfURL = newPdfURL
	}

	timelineEntity := &entity.Timeline{
		ID:          input.ID,
		Period:      input.Period,
		PdfURL:      pdfURL,
		IsHighlight: input.IsHighlight,
		UpdatedAt:   time.Now(),
	}

	if err := timelineEntity.Validate(); err != nil {

		if oldPdfURL != "" {
			_ = uc.StorageService.DeletePDF(pdfURL, "timelines")
		}
		return nil, err
	}

	timelineUpdated, err := uc.Repository.UpdateTimelineByID(timelineEntity)
	if err != nil {

		if oldPdfURL != "" {
			_ = uc.StorageService.DeletePDF(pdfURL, "timelines")
		}
		return nil, err
	}

	if oldPdfURL != "" {
		_ = uc.StorageService.DeletePDF(oldPdfURL, "timelines")
	}

	return &UpdateTimelineByIDOutputDTO{
		ID:          timelineUpdated.ID,
		Period:      timelineUpdated.Period,
		PdfURL:      timelineUpdated.PdfURL,
		IsHighlight: timelineUpdated.IsHighlight,
		CreatedAt:   timelineUpdated.CreatedAt,
		UpdatedAt:   timelineUpdated.UpdatedAt,
	}, nil
}
