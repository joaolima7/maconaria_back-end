package worker_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/worker"
)

type GetAllWorkersOutputDTO struct {
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

type GetAllWorkersUseCase struct {
	Repository worker.GetAllWorkersRepository
}

func NewGetAllWorkersUseCase(repository worker.GetAllWorkersRepository) *GetAllWorkersUseCase {
	return &GetAllWorkersUseCase{
		Repository: repository,
	}
}

func (uc *GetAllWorkersUseCase) Execute() ([]*GetAllWorkersOutputDTO, error) {
	workers, err := uc.Repository.GetAllWorkers()
	if err != nil {
		return nil, err
	}

	output := make([]*GetAllWorkersOutputDTO, len(workers))
	for i, w := range workers {

		imageData := ""
		if len(w.ImageData) > 0 {
			imageData = base64.StdEncoding.EncodeToString(w.ImageData)
		}

		output[i] = &GetAllWorkersOutputDTO{
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
			ImageData:         imageData,
			Deceased:          w.Deceased,
			CreatedAt:         w.CreatedAt,
			UpdatedAt:         w.UpdatedAt,
		}
	}

	return output, nil
}
