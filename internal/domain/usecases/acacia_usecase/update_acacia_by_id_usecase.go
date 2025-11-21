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

type UpdateAcaciaByIDInputDTO struct {
	ID          string   `json:"id"`
	Name        string   `json:"name" validate:"required,min=3"`
	Terms       []string `json:"terms"`
	IsPresident bool     `json:"is_president"`
	Deceased    bool     `json:"deceased"`
	ImageData   string   `json:"image_data,omitempty"`
}

type UpdateAcaciaByIDOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Terms       []string  `json:"terms"`
	IsPresident bool      `json:"is_president"`
	Deceased    bool      `json:"deceased"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateAcaciaByIDUseCase struct {
	Repository     acacia.UpdateAcaciaByIDRepository
	GetRepository  acacia.GetAcaciaByIDRepository
	StorageService storage.StorageService
}

func NewUpdateAcaciaByIDUseCase(
	repository acacia.UpdateAcaciaByIDRepository,
	getRepository acacia.GetAcaciaByIDRepository,
	storageService storage.StorageService,
) *UpdateAcaciaByIDUseCase {
	return &UpdateAcaciaByIDUseCase{
		Repository:     repository,
		GetRepository:  getRepository,
		StorageService: storageService,
	}
}

func (uc *UpdateAcaciaByIDUseCase) Execute(input UpdateAcaciaByIDInputDTO) (*UpdateAcaciaByIDOutputDTO, error) {

	existingAcacia, err := uc.GetRepository.GetAcaciaByID(input.ID)
	if err != nil {
		return nil, err
	}

	if input.IsPresident && len(input.Terms) == 0 {
		return nil, apperrors.NewValidationError("mandatos", "Os períodos são obrigatórios para presidentes!")
	}

	imageURL := existingAcacia.ImageURL
	var oldImageURL string

	if input.ImageData != "" {
		imageData, err := base64.StdEncoding.DecodeString(input.ImageData)
		if err != nil {
			return nil, apperrors.NewValidationError("imagem", "Imagem em formato inválido!")
		}

		filename := fmt.Sprintf("acacia_%s_%s.jpg", input.ID, uuid.New().String())

		newImageURL, err := uc.StorageService.UploadImage(imageData, filename, "acacias")
		if err != nil {
			return nil, apperrors.NewInternalError("Erro ao fazer upload da imagem", err)
		}

		oldImageURL = imageURL
		imageURL = newImageURL
	}

	acaciaEntity := &entity.Acacia{
		ID:          input.ID,
		Name:        input.Name,
		Terms:       input.Terms,
		IsPresident: input.IsPresident,
		Deceased:    input.Deceased,
		ImageURL:    imageURL,
		UpdatedAt:   time.Now(),
	}

	if err := acaciaEntity.Validate(); err != nil {

		if oldImageURL != "" {
			_ = uc.StorageService.DeleteImage(imageURL, "acacias")
		}
		return nil, err
	}

	acaciaUpdated, err := uc.Repository.UpdateAcaciaByID(acaciaEntity)
	if err != nil {

		if oldImageURL != "" {
			_ = uc.StorageService.DeleteImage(imageURL, "acacias")
		}
		return nil, err
	}

	if oldImageURL != "" {
		_ = uc.StorageService.DeleteImage(oldImageURL, "acacias")
	}

	return &UpdateAcaciaByIDOutputDTO{
		ID:          acaciaUpdated.ID,
		Name:        acaciaUpdated.Name,
		Terms:       acaciaUpdated.Terms,
		IsPresident: acaciaUpdated.IsPresident,
		Deceased:    acaciaUpdated.Deceased,
		ImageURL:    acaciaUpdated.ImageURL,
		CreatedAt:   acaciaUpdated.CreatedAt,
		UpdatedAt:   acaciaUpdated.UpdatedAt,
	}, nil
}
