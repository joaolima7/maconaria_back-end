package library_usecase

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"
)

type CreateLibraryInputDTO struct {
	Title            string `json:"title" validate:"required"`
	SmallDescription string `json:"small_description" validate:"required"`
	Degree           string `json:"degree" validate:"required,oneof=apprentice companion master"`
	FileData         string `json:"file_data"`
	CoverData        string `json:"cover_data"`
	Link             string `json:"link"`
}

type CreateLibraryOutputDTO struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	SmallDescription string    `json:"small_description"`
	Degree           string    `json:"degree"`
	FileData         string    `json:"file_data"`
	CoverData        string    `json:"cover_data"`
	Link             string    `json:"link"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreateLibraryUseCase struct {
	Repository library.CreateLibraryRepository
}

func NewCreateLibraryUseCase(repository library.CreateLibraryRepository) *CreateLibraryUseCase {
	return &CreateLibraryUseCase{
		Repository: repository,
	}
}

func (uc *CreateLibraryUseCase) Execute(input CreateLibraryInputDTO) (*CreateLibraryOutputDTO, error) {
	var fileData []byte
	var coverData []byte
	var err error

	if input.FileData != "" {
		fileData, err = base64.StdEncoding.DecodeString(input.FileData)
		if err != nil {
			return nil, apperrors.NewValidationError("arquivo", "Arquivo em formato inválido!")
		}
	}

	if input.CoverData != "" {
		coverData, err = base64.StdEncoding.DecodeString(input.CoverData)
		if err != nil {
			return nil, apperrors.NewValidationError("capa", "Capa em formato inválido!")
		}
	}

	libraryEntity, err := entity.NewLibrary(
		uuid.New().String(),
		input.Title,
		input.SmallDescription,
		entity.UserDegree(input.Degree),
		fileData,
		coverData,
		input.Link,
	)
	if err != nil {
		return nil, err
	}

	libraryCreated, err := uc.Repository.CreateLibrary(libraryEntity)
	if err != nil {
		return nil, err
	}

	fileDataBase64 := ""
	if len(libraryCreated.FileData) > 0 {
		fileDataBase64 = base64.StdEncoding.EncodeToString(libraryCreated.FileData)
	}

	coverDataBase64 := ""
	if len(libraryCreated.CoverData) > 0 {
		coverDataBase64 = base64.StdEncoding.EncodeToString(libraryCreated.CoverData)
	}

	return &CreateLibraryOutputDTO{
		ID:               libraryCreated.ID,
		Title:            libraryCreated.Title,
		SmallDescription: libraryCreated.SmallDescription,
		Degree:           string(libraryCreated.Degree),
		FileData:         fileDataBase64,
		CoverData:        coverDataBase64,
		Link:             libraryCreated.Link,
		CreatedAt:        libraryCreated.CreatedAt,
		UpdatedAt:        libraryCreated.UpdatedAt,
	}, nil
}
