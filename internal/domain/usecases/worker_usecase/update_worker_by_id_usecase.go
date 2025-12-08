package worker_usecase

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/worker"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type UpdateWorkerByIDInputDTO struct {
	ID                string  `json:"id"`
	Number            int32   `json:"number" validate:"required,gt=0"`
	Name              string  `json:"name" validate:"required,min=3"`
	Registration      string  `json:"registration" validate:"required"`
	BirthDate         string  `json:"birth_date" validate:"required"`
	InitiationDate    *string `json:"initiation_date,omitempty"`
	ElevationDate     *string `json:"elevation_date,omitempty"`
	ExaltationDate    *string `json:"exaltation_date,omitempty"`
	AffiliationDate   *string `json:"affiliation_date,omitempty"`
	InstallationDate  *string `json:"installation_date,omitempty"`
	EmeritusMasonDate *string `json:"emeritus_mason_date,omitempty"`
	ProvectMasonDate  *string `json:"provect_mason_date,omitempty"`
	ImageData         string  `json:"image_data,omitempty"`
	Deceased          bool    `json:"deceased"`
}

type UpdateWorkerByIDOutputDTO struct {
	ID                string     `json:"id"`
	Number            int32      `json:"number"`
	Name              string     `json:"name"`
	Registration      string     `json:"registration"`
	BirthDate         time.Time  `json:"birth_date"`
	InitiationDate    *time.Time `json:"initiation_date,omitempty"`
	ElevationDate     *time.Time `json:"elevation_date,omitempty"`
	ExaltationDate    *time.Time `json:"exaltation_date,omitempty"`
	AffiliationDate   *time.Time `json:"affiliation_date,omitempty"`
	InstallationDate  *time.Time `json:"installation_date,omitempty"`
	EmeritusMasonDate *time.Time `json:"emeritus_mason_date,omitempty"`
	ProvectMasonDate  *time.Time `json:"provect_mason_date,omitempty"`
	ImageURL          string     `json:"image_url"`
	Deceased          bool       `json:"deceased"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type UpdateWorkerByIDUseCase struct {
	Repository     worker.UpdateWorkerByIDRepository
	GetRepository  worker.GetWorkerByIDRepository
	StorageService storage.StorageService
}

func NewUpdateWorkerByIDUseCase(
	repository worker.UpdateWorkerByIDRepository,
	getRepository worker.GetWorkerByIDRepository,
	storageService storage.StorageService,
) *UpdateWorkerByIDUseCase {
	return &UpdateWorkerByIDUseCase{
		Repository:     repository,
		GetRepository:  getRepository,
		StorageService: storageService,
	}
}

func (uc *UpdateWorkerByIDUseCase) Execute(input UpdateWorkerByIDInputDTO) (*UpdateWorkerByIDOutputDTO, error) {

	existingWorker, err := uc.GetRepository.GetWorkerByID(input.ID)
	if err != nil {
		return nil, err
	}

	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return nil, apperrors.NewValidationError("data de nascimento", "Formato da data de nascimento inválido!")
	}

	var initiationDate, elevationDate, exaltationDate, affiliationDate, installationDate *time.Time
	var emeritusMasonDate, provectMasonDate *time.Time

	if input.InitiationDate != nil {
		date, err := time.Parse("2006-01-02", *input.InitiationDate)
		if err != nil {
			return nil, apperrors.NewValidationError("data de iniciação", "Formato da data de iniciação inválido!")
		}
		initiationDate = &date
	}

	if input.ElevationDate != nil {
		date, err := time.Parse("2006-01-02", *input.ElevationDate)
		if err != nil {
			return nil, apperrors.NewValidationError("data de elevação", "Formato da data de elevação inválido!")
		}
		elevationDate = &date
	}

	if input.ExaltationDate != nil {
		date, err := time.Parse("2006-01-02", *input.ExaltationDate)
		if err != nil {
			return nil, apperrors.NewValidationError("data de exaltação", "Formato da data de exaltação inválido!")
		}
		exaltationDate = &date
	}

	if input.AffiliationDate != nil {
		date, err := time.Parse("2006-01-02", *input.AffiliationDate)
		if err != nil {
			return nil, apperrors.NewValidationError("data de afiliação", "Formato da data de afiliação inválido!")
		}
		affiliationDate = &date
	}

	if input.InstallationDate != nil {
		date, err := time.Parse("2006-01-02", *input.InstallationDate)
		if err != nil {
			return nil, apperrors.NewValidationError("data de instalação", "Formato da data de instalação inválido!")
		}
		installationDate = &date
	}
	if input.EmeritusMasonDate != nil {
		date, err := time.Parse("2006-01-02", *input.EmeritusMasonDate)
		if err != nil {
			return nil, apperrors.NewValidationError("data de maçom emérito", "Formato da data de maçom emérito inválido!")
		}
		emeritusMasonDate = &date
	}
	if input.ProvectMasonDate != nil {
		date, err := time.Parse("2006-01-02", *input.ProvectMasonDate)
		if err != nil {
			return nil, apperrors.NewValidationError("data de maçom provecto", "Formato da data de maçom provecto inválido!")
		}
		provectMasonDate = &date
	}

	imageURL := existingWorker.ImageURL
	var oldImageURL string

	if input.ImageData != "" {
		imageData, err := base64.StdEncoding.DecodeString(input.ImageData)
		if err != nil {
			return nil, apperrors.NewValidationError("imagem", "Imagem em formato inválido!")
		}

		filename := fmt.Sprintf("worker_%s_%s.jpg", input.ID, uuid.New().String())

		newImageURL, err := uc.StorageService.UploadImage(imageData, filename, "workers")
		if err != nil {
			return nil, apperrors.NewInternalError("Erro ao fazer upload da imagem", err)
		}

		oldImageURL = imageURL
		imageURL = newImageURL
	}

	workerEntity := &entity.Worker{
		ID:                input.ID,
		Number:            input.Number,
		Name:              input.Name,
		Registration:      input.Registration,
		BirthDate:         birthDate,
		InitiationDate:    initiationDate,
		ElevationDate:     elevationDate,
		ExaltationDate:    exaltationDate,
		AffiliationDate:   affiliationDate,
		InstallationDate:  installationDate,
		EmeritusMasonDate: emeritusMasonDate,
		ProvectMasonDate:  provectMasonDate,
		ImageURL:          imageURL,
		Deceased:          input.Deceased,
		UpdatedAt:         time.Now(),
	}

	if err := workerEntity.Validate(); err != nil {

		if oldImageURL != "" {
			_ = uc.StorageService.DeleteImage(imageURL, "workers")
		}
		return nil, err
	}

	workerUpdated, err := uc.Repository.UpdateWorkerByID(workerEntity)
	if err != nil {

		if oldImageURL != "" {
			_ = uc.StorageService.DeleteImage(imageURL, "workers")
		}
		return nil, err
	}

	if oldImageURL != "" {
		_ = uc.StorageService.DeleteImage(oldImageURL, "workers")
	}

	return &UpdateWorkerByIDOutputDTO{
		ID:                workerUpdated.ID,
		Number:            workerUpdated.Number,
		Name:              workerUpdated.Name,
		Registration:      workerUpdated.Registration,
		BirthDate:         workerUpdated.BirthDate,
		InitiationDate:    workerUpdated.InitiationDate,
		ElevationDate:     workerUpdated.ElevationDate,
		ExaltationDate:    workerUpdated.ExaltationDate,
		AffiliationDate:   workerUpdated.AffiliationDate,
		InstallationDate:  workerUpdated.InstallationDate,
		EmeritusMasonDate: workerUpdated.EmeritusMasonDate,
		ProvectMasonDate:  workerUpdated.ProvectMasonDate,
		ImageURL:          workerUpdated.ImageURL,
		Deceased:          workerUpdated.Deceased,
		CreatedAt:         workerUpdated.CreatedAt,
		UpdatedAt:         workerUpdated.UpdatedAt,
	}, nil
}
