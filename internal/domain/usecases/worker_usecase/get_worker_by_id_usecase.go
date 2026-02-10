package worker_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/worker"
)

type GetWorkerByIDInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type GetWorkerByIDOutputDTO struct {
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
	Terms             []string   `json:"terms"`
	IsActive          bool       `json:"is_active"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type GetWorkerByIDUseCase struct {
	Repository worker.GetWorkerByIDRepository
}

func NewGetWorkerByIDUseCase(repository worker.GetWorkerByIDRepository) *GetWorkerByIDUseCase {
	return &GetWorkerByIDUseCase{
		Repository: repository,
	}
}

func (uc *GetWorkerByIDUseCase) Execute(input GetWorkerByIDInputDTO) (*GetWorkerByIDOutputDTO, error) {
	w, err := uc.Repository.GetWorkerByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetWorkerByIDOutputDTO{
		ID:                w.ID,
		Number:            w.Number,
		Name:              w.Name,
		Registration:      w.Registration,
		BirthDate:         w.BirthDate,
		InitiationDate:    w.InitiationDate,
		ElevationDate:     w.ElevationDate,
		ExaltationDate:    w.ExaltationDate,
		AffiliationDate:   w.AffiliationDate,
		InstallationDate:  w.InstallationDate,
		EmeritusMasonDate: w.EmeritusMasonDate,
		ProvectMasonDate:  w.ProvectMasonDate,
		ImageURL:          w.ImageURL,
		Deceased:          w.Deceased,
		IsPresident:       w.IsPresident,
		Terms:             w.Terms,
		IsActive:          w.IsActive,
		CreatedAt:         w.CreatedAt,
		UpdatedAt:         w.UpdatedAt,
	}, nil
}
