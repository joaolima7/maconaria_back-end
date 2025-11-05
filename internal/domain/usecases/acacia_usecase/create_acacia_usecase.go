package acacia_usecase

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/acacia"
)

type CreateAcaciaInputDTO struct {
	Name        string   `json:"name" validate:"required,min=3"`
	Terms       []string `json:"terms"`
	IsPresident bool     `json:"is_president"`
	Deceased    bool     `json:"deceased"`
	ImageData   string   `json:"image_data" validate:"required"`
}

type CreateAcaciaOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Terms       []string  `json:"terms"`
	IsPresident bool      `json:"is_president"`
	Deceased    bool      `json:"deceased"`
	ImageData   string    `json:"image_data"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateAcaciaUseCase struct {
	Repository acacia.CreateAcaciaRepository
}

func NewCreateAcaciaUseCase(repository acacia.CreateAcaciaRepository) *CreateAcaciaUseCase {
	return &CreateAcaciaUseCase{
		Repository: repository,
	}
}

func (uc *CreateAcaciaUseCase) Execute(input CreateAcaciaInputDTO) (*CreateAcaciaOutputDTO, error) {
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

	acaciaEntity, err := entity.NewAcacia(
		uuid.New().String(),
		input.Name,
		input.Terms,
		input.IsPresident,
		input.Deceased,
		imageData,
	)
	if err != nil {
		return nil, err
	}

	acaciaCreated, err := uc.Repository.CreateAcacia(acaciaEntity)
	if err != nil {
		return nil, err
	}

	imageDataBase64 := ""
	if len(acaciaCreated.ImageData) > 0 {
		imageDataBase64 = base64.StdEncoding.EncodeToString(acaciaCreated.ImageData)
	}

	return &CreateAcaciaOutputDTO{
		ID:          acaciaCreated.ID,
		Name:        acaciaCreated.Name,
		Terms:       acaciaCreated.Terms,
		IsPresident: acaciaCreated.IsPresident,
		Deceased:    acaciaCreated.Deceased,
		ImageData:   imageDataBase64,
		CreatedAt:   acaciaCreated.CreatedAt,
		UpdatedAt:   acaciaCreated.UpdatedAt,
	}, nil
}
