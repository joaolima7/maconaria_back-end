package worker_usecase

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/worker"
)

type CreateWorkerInputDTO struct {
	Number            int32   `json:"number" validate:"required,gt=0"`
	Name              string  `json:"name" validate:"required,min=3"`
	Registration      string  `json:"registration" validate:"required"`
	BirthDate         string  `json:"birth_date" validate:"required"`
	InitiationDate    string  `json:"initiation_date" validate:"required"`
	ElevationDate     string  `json:"elevation_date" validate:"required"`
	ExaltationDate    string  `json:"exaltation_date" validate:"required"`
	AffiliationDate   string  `json:"affiliation_date" validate:"required"`
	InstallationDate  string  `json:"installation_date" validate:"required"`
	EmeritusMasonDate *string `json:"emeritus_mason_date,omitempty"`
	ProvectMasonDate  *string `json:"provect_mason_date,omitempty"`
	ImageData         string  `json:"image_data"`
	Deceased          bool    `json:"deceased"`
}

type CreateWorkerOutputDTO struct {
	ID                string     `json:"id"`
	Number            int32      `json:"number"`
	Name              string     `json:"name"`
	Registration      string     `json:"registration"`
	BirthDate         time.Time  `json:"birth_date"`
	InitiationDate    time.Time  `json:"initiation_date"`
	ElevationDate     time.Time  `json:"elevation_date"`
	ExaltationDate    time.Time  `json:"exaltation_date"`
	AffiliationDate   time.Time  `json:"affiliation_date"`
	InstallationDate  time.Time  `json:"installation_date"`
	EmeritusMasonDate *time.Time `json:"emeritus_mason_date,omitempty"`
	ProvectMasonDate  *time.Time `json:"provect_mason_date,omitempty"`
	ImageData         string     `json:"image_data"`
	Deceased          bool       `json:"deceased"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type CreateWorkerUseCase struct {
	Repository worker.CreateWorkerRepository
}

func NewCreateWorkerUseCase(repository worker.CreateWorkerRepository) *CreateWorkerUseCase {
	return &CreateWorkerUseCase{
		Repository: repository,
	}
}

func (uc *CreateWorkerUseCase) Execute(input CreateWorkerInputDTO) (*CreateWorkerOutputDTO, error) {
	if input.ImageData == "" {
		return nil, apperrors.NewValidationError("imagem", "A imagem é obrigatória!")
	}

	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return nil, apperrors.NewValidationError("data de nascimento", "Formato da data de nascimento inválido!")
	}

	initiationDate, err := time.Parse("2006-01-02", input.InitiationDate)
	if err != nil {
		return nil, apperrors.NewValidationError("data de iniciação", "Formato da data de iniciação inválido!")
	}

	elevationDate, err := time.Parse("2006-01-02", input.ElevationDate)
	if err != nil {
		return nil, apperrors.NewValidationError("data de elevação", "Formato da data de elevação inválido!")
	}

	exaltationDate, err := time.Parse("2006-01-02", input.ExaltationDate)
	if err != nil {
		return nil, apperrors.NewValidationError("data de exaltação", "Formato da data de exaltação inválido!")
	}

	affiliationDate, err := time.Parse("2006-01-02", input.AffiliationDate)
	if err != nil {
		return nil, apperrors.NewValidationError("data de afiliação", "Formato da data de afiliação inválido!")
	}

	installationDate, err := time.Parse("2006-01-02", input.InstallationDate)
	if err != nil {
		return nil, apperrors.NewValidationError("data de instalação", "Formato da data de instalação inválido!")
	}

	var emeritusMasonDate, provectMasonDate *time.Time
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

	workerEntity, err := entity.NewWorker(
		uuid.New().String(),
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
		imageData,
		input.Deceased,
	)
	if err != nil {
		return nil, err
	}

	workerCreated, err := uc.Repository.CreateWorker(workerEntity)
	if err != nil {
		return nil, err
	}

	imageDataBase64 := ""
	if len(workerCreated.ImageData) > 0 {
		imageDataBase64 = base64.StdEncoding.EncodeToString(workerCreated.ImageData)
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
		ImageData:         imageDataBase64,
		Deceased:          workerCreated.Deceased,
		CreatedAt:         workerCreated.CreatedAt,
		UpdatedAt:         workerCreated.UpdatedAt,
	}, nil
}
