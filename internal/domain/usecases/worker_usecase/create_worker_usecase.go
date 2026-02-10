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

type CreateWorkerInputDTO struct {
	Number            int32    `json:"number" validate:"required,gt=0"`
	Name              string   `json:"name" validate:"required,min=3"`
	Registration      string   `json:"registration" validate:"required"`
	BirthDate         string   `json:"birth_date" validate:"required"`
	InitiationDate    *string  `json:"initiation_date,omitempty"`
	ElevationDate     *string  `json:"elevation_date,omitempty"`
	ExaltationDate    *string  `json:"exaltation_date,omitempty"`
	AffiliationDate   *string  `json:"affiliation_date,omitempty"`
	InstallationDate  *string  `json:"installation_date,omitempty"`
	EmeritusMasonDate *string  `json:"emeritus_mason_date,omitempty"`
	ProvectMasonDate  *string  `json:"provect_mason_date,omitempty"`
	ImageData         string   `json:"image_data" validate:"required,base64"`
	Deceased          bool     `json:"deceased"`
	IsPresident       bool     `json:"is_president"`
	Terms             []string `json:"terms"`
	IsActive          bool     `json:"is_active"`
}

type CreateWorkerOutputDTO struct {
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
	IsPresident       bool       `json:"is_president"`
	Terms             []string   `json:"terms,omitempty"`
	IsActive          bool       `json:"is_active"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type CreateWorkerUseCase struct {
	Repository     worker.CreateWorkerRepository
	StorageService storage.StorageService
}

func NewCreateWorkerUseCase(
	repository worker.CreateWorkerRepository,
	storageService storage.StorageService,
) *CreateWorkerUseCase {
	return &CreateWorkerUseCase{
		Repository:     repository,
		StorageService: storageService,
	}
}

func (uc *CreateWorkerUseCase) Execute(input CreateWorkerInputDTO) (*CreateWorkerOutputDTO, error) {
	if input.ImageData == "" {
		return nil, apperrors.NewValidationError("imagem", "A imagem é obrigatória!")
	}

	if input.IsPresident && len(input.Terms) == 0 {
		return nil, apperrors.NewValidationError("mandatos", "Os períodos são obrigatórios para presidentes!")
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

	imageData, err := base64.StdEncoding.DecodeString(input.ImageData)
	if err != nil {
		return nil, apperrors.NewValidationError("imagem", "Imagem em formato inválido!")
	}

	workerID := uuid.New().String()
	filename := fmt.Sprintf("worker_%s_%s.jpg", workerID, uuid.New().String())

	imageURL, err := uc.StorageService.UploadImage(imageData, filename, "workers")
	if err != nil {
		return nil, apperrors.NewInternalError("Erro ao fazer upload da imagem", err)
	}

	workerEntity, err := entity.NewWorker(
		workerID,
		input.Number,
		input.Name,
		input.Registration,
		birthDate,
		initiationDate,
		elevationDate,
		exaltationDate,
		affiliationDate,
		installationDate,
		emeritusMasonDate,
		provectMasonDate,
		imageURL,
		input.Deceased,
		input.IsPresident,
		input.Terms,
		input.IsActive,
	)
	if err != nil {

		_ = uc.StorageService.DeleteImage(imageURL, "workers")
		return nil, err
	}

	workerCreated, err := uc.Repository.CreateWorker(workerEntity)
	if err != nil {

		_ = uc.StorageService.DeleteImage(imageURL, "workers")
		return nil, err
	}

	return &CreateWorkerOutputDTO{
		ID:                workerCreated.ID,
		Number:            workerCreated.Number,
		Name:              workerCreated.Name,
		Registration:      workerCreated.Registration,
		BirthDate:         workerCreated.BirthDate,
		InitiationDate:    workerCreated.InitiationDate,
		ElevationDate:     workerCreated.ElevationDate,
		ExaltationDate:    workerCreated.ExaltationDate,
		AffiliationDate:   workerCreated.AffiliationDate,
		InstallationDate:  workerCreated.InstallationDate,
		EmeritusMasonDate: workerCreated.EmeritusMasonDate,
		ProvectMasonDate:  workerCreated.ProvectMasonDate,
		ImageURL:          workerCreated.ImageURL,
		Deceased:          workerCreated.Deceased,
		IsPresident:       workerCreated.IsPresident,
		Terms:             workerCreated.Terms,
		IsActive:          workerCreated.IsActive,
		CreatedAt:         workerCreated.CreatedAt,
		UpdatedAt:         workerCreated.UpdatedAt,
	}, nil
}
