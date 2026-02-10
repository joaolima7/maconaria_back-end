package acacia_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/acacia"
)

type GetAcaciaByIDInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type GetAcaciaByIDOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Terms       []string  `json:"terms"`
	IsPresident bool      `json:"is_president"`
	Deceased    bool      `json:"deceased"`
	IsActive    bool      `json:"is_active"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAcaciaByIDUseCase struct {
	Repository acacia.GetAcaciaByIDRepository
}

func NewGetAcaciaByIDUseCase(repository acacia.GetAcaciaByIDRepository) *GetAcaciaByIDUseCase {
	return &GetAcaciaByIDUseCase{
		Repository: repository,
	}
}

func (uc *GetAcaciaByIDUseCase) Execute(input GetAcaciaByIDInputDTO) (*GetAcaciaByIDOutputDTO, error) {
	a, err := uc.Repository.GetAcaciaByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetAcaciaByIDOutputDTO{
		ID:          a.ID,
		Name:        a.Name,
		Terms:       a.Terms,
		IsPresident: a.IsPresident,
		Deceased:    a.Deceased,
		IsActive:    a.IsActive,
		ImageURL:    a.ImageURL,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}, nil
}
