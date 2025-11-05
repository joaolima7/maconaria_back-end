package acacia_usecase

import (
	"encoding/base64"
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
	ImageData   string    `json:"image_data"`
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

	imageData := ""
	if len(a.ImageData) > 0 {
		imageData = base64.StdEncoding.EncodeToString(a.ImageData)
	}

	return &GetAcaciaByIDOutputDTO{
		ID:          a.ID,
		Name:        a.Name,
		Terms:       a.Terms,
		IsPresident: a.IsPresident,
		Deceased:    a.Deceased,
		ImageData:   imageData,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}, nil
}
