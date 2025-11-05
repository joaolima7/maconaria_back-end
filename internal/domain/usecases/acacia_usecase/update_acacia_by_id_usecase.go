package acacia_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/acacia"
)

type UpdateAcaciaByIDInputDTO struct {
	ID          string   `json:"id"`
	Name        string   `json:"name" validate:"required,min=3"`
	Terms       []string `json:"terms"`
	IsPresident bool     `json:"is_president"`
	Deceased    bool     `json:"deceased"`
	ImageData   string   `json:"image_data" validate:"required"`
}

type UpdateAcaciaByIDOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Terms       []string  `json:"terms"`
	IsPresident bool      `json:"is_president"`
	Deceased    bool      `json:"deceased"`
	ImageData   string    `json:"image_data"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateAcaciaByIDUseCase struct {
	Repository acacia.UpdateAcaciaByIDRepository
}

func NewUpdateAcaciaByIDUseCase(repository acacia.UpdateAcaciaByIDRepository) *UpdateAcaciaByIDUseCase {
	return &UpdateAcaciaByIDUseCase{
		Repository: repository,
	}
}

func (uc *UpdateAcaciaByIDUseCase) Execute(input UpdateAcaciaByIDInputDTO) (*UpdateAcaciaByIDOutputDTO, error) {
	if input.ImageData == "" {
		return nil, apperrors.NewValidationError("imagem", "A imagem é obrigatória!")
	}

	if input.IsPresident && len(input.Terms) == 0 {
		return nil, apperrors.NewValidationError("mandatos", "Os períodos são obrigatórios para presidentes!")
	}

	imageData, err := base64.StdEncoding.DecodeString(input.ImageData)
	if err != nil {
		return nil, apperrors.NewValidationError("imagem", "Imagem em formato inválido!")
	}

	acaciaEntity := &entity.Acacia{
		ID:          input.ID,
		Name:        input.Name,
		Terms:       input.Terms,
		IsPresident: input.IsPresident,
		Deceased:    input.Deceased,
		ImageData:   imageData,
		UpdatedAt:   time.Now(),
	}

	if err := acaciaEntity.Validate(); err != nil {
		return nil, err
	}

	acaciaUpdated, err := uc.Repository.UpdateAcaciaByID(acaciaEntity)
	if err != nil {
		return nil, err
	}

	imageDataBase64 := ""
	if len(acaciaUpdated.ImageData) > 0 {
		imageDataBase64 = base64.StdEncoding.EncodeToString(acaciaUpdated.ImageData)
	}

	return &UpdateAcaciaByIDOutputDTO{
		ID:          acaciaUpdated.ID,
		Name:        acaciaUpdated.Name,
		Terms:       acaciaUpdated.Terms,
		IsPresident: acaciaUpdated.IsPresident,
		Deceased:    acaciaUpdated.Deceased,
		ImageData:   imageDataBase64,
		CreatedAt:   acaciaUpdated.CreatedAt,
		UpdatedAt:   acaciaUpdated.UpdatedAt,
	}, nil
}
