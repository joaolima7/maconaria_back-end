package acacia_usecase

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/acacia"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type CreateAcaciaInputDTO struct {
	Name        string   `json:"name" validate:"required,min=3"`
	Terms       []string `json:"terms"`
	IsPresident bool     `json:"is_president"`
	Deceased    bool     `json:"deceased"`
	IsActive    bool     `json:"is_active"`
	ImageData   string   `json:"image_data" validate:"required,base64"`
}

type CreateAcaciaOutputDTO struct {
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

type CreateAcaciaUseCase struct {
	Repository     acacia.CreateAcaciaRepository
	StorageService storage.StorageService
}

func NewCreateAcaciaUseCase(
	repository acacia.CreateAcaciaRepository,
	storageService storage.StorageService,
) *CreateAcaciaUseCase {
	return &CreateAcaciaUseCase{
		Repository:     repository,
		StorageService: storageService,
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

	acaciaID := uuid.New().String()
	filename := fmt.Sprintf("acacia_%s_%s.jpg", acaciaID, uuid.New().String())

	imageURL, err := uc.StorageService.UploadImage(imageData, filename, "acacias")
	if err != nil {
		return nil, apperrors.NewInternalError("Erro ao fazer upload da imagem", err)
	}

	acaciaEntity, err := entity.NewAcacia(
		acaciaID,
		input.Name,
		input.Terms,
		input.IsPresident,
		input.Deceased,
		imageURL,
		input.IsActive,
	)
	if err != nil {

		_ = uc.StorageService.DeleteImage(imageURL, "acacias")
		return nil, err
	}

	acaciaCreated, err := uc.Repository.CreateAcacia(acaciaEntity)
	if err != nil {

		_ = uc.StorageService.DeleteImage(imageURL, "acacias")
		return nil, err
	}

	return &CreateAcaciaOutputDTO{
		ID:          acaciaCreated.ID,
		Name:        acaciaCreated.Name,
		Terms:       acaciaCreated.Terms,
		IsPresident: acaciaCreated.IsPresident,
		Deceased:    acaciaCreated.Deceased,
		IsActive:    acaciaCreated.IsActive,
		ImageURL:    acaciaCreated.ImageURL,
		CreatedAt:   acaciaCreated.CreatedAt,
		UpdatedAt:   acaciaCreated.UpdatedAt,
	}, nil
}
