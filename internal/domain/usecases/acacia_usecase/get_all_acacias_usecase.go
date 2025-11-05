package acacia_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/acacia"
)

type GetAllAcaciasOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Terms       []string  `json:"terms"`
	IsPresident bool      `json:"is_president"`
	Deceased    bool      `json:"deceased"`
	ImageData   string    `json:"image_data"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllAcaciasUseCase struct {
	Repository acacia.GetAllAcaciasRepository
}

func NewGetAllAcaciasUseCase(repository acacia.GetAllAcaciasRepository) *GetAllAcaciasUseCase {
	return &GetAllAcaciasUseCase{
		Repository: repository,
	}
}

func (uc *GetAllAcaciasUseCase) Execute() ([]*GetAllAcaciasOutputDTO, error) {
	acacias, err := uc.Repository.GetAllAcacias()
	if err != nil {
		return nil, err
	}

	output := make([]*GetAllAcaciasOutputDTO, len(acacias))
	for i, a := range acacias {
		imageData := ""
		if len(a.ImageData) > 0 {
			imageData = base64.StdEncoding.EncodeToString(a.ImageData)
		}

		output[i] = &GetAllAcaciasOutputDTO{
			ID:          a.ID,
			Name:        a.Name,
			Terms:       a.Terms,
			IsPresident: a.IsPresident,
			Deceased:    a.Deceased,
			ImageData:   imageData,
			CreatedAt:   a.CreatedAt,
			UpdatedAt:   a.UpdatedAt,
		}
	}

	return output, nil
}
