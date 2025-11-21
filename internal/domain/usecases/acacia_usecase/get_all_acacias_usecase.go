package acacia_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/acacia"
)

type GetAllAcaciasOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Terms       []string  `json:"terms"`
	IsPresident bool      `json:"is_president"`
	Deceased    bool      `json:"deceased"`
	ImageURL    string    `json:"image_url"`
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
		output[i] = &GetAllAcaciasOutputDTO{
			ID:          a.ID,
			Name:        a.Name,
			Terms:       a.Terms,
			IsPresident: a.IsPresident,
			Deceased:    a.Deceased,
			ImageURL:    a.ImageURL,
			CreatedAt:   a.CreatedAt,
			UpdatedAt:   a.UpdatedAt,
		}
	}

	return output, nil
}
