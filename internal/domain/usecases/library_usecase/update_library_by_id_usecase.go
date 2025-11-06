package library_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"
)

type UpdateLibraryByIDInputDTO struct {
	ID               string `json:"id"`
	Title            string `json:"title" validate:"required"`
	SmallDescription string `json:"small_description" validate:"required"`
	Degree           string `json:"degree" validate:"required,oneof=apprentice companion master"`
	FileData         string `json:"file_data"`
	CoverData        string `json:"cover_data"`
	Link             string `json:"link"`
}

type UpdateLibraryByIDOutputDTO struct {
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

type UpdateLibraryByIDUseCase struct {
	Repository library.UpdateLibraryByIDRepository
}

func NewUpdateLibraryByIDUseCase(repository library.UpdateLibraryByIDRepository) *UpdateLibraryByIDUseCase {
	return &UpdateLibraryByIDUseCase{
		Repository: repository,
	}
}

func (uc *UpdateLibraryByIDUseCase) Execute(input UpdateLibraryByIDInputDTO) (*UpdateLibraryByIDOutputDTO, error) {
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

	libraryEntity := &entity.Library{
		ID:               input.ID,
		Title:            input.Title,
		SmallDescription: input.SmallDescription,
		Degree:           entity.UserDegree(input.Degree),
		FileData:         fileData,
		CoverData:        coverData,
		Link:             input.Link,
		UpdatedAt:        time.Now(),
	}

	if err := libraryEntity.Validate(); err != nil {
		return nil, err
	}

	libraryUpdated, err := uc.Repository.UpdateLibraryByID(libraryEntity)
	if err != nil {
		return nil, err
	}

	fileDataBase64 := ""
	if len(libraryUpdated.FileData) > 0 {
		fileDataBase64 = base64.StdEncoding.EncodeToString(libraryUpdated.FileData)
	}

	coverDataBase64 := ""
	if len(libraryUpdated.CoverData) > 0 {
		coverDataBase64 = base64.StdEncoding.EncodeToString(libraryUpdated.CoverData)
	}

	return &UpdateLibraryByIDOutputDTO{
		ID:               libraryUpdated.ID,
		Title:            libraryUpdated.Title,
		SmallDescription: libraryUpdated.SmallDescription,
		Degree:           string(libraryUpdated.Degree),
		FileData:         fileDataBase64,
		CoverData:        coverDataBase64,
		Link:             libraryUpdated.Link,
		CreatedAt:        libraryUpdated.CreatedAt,
		UpdatedAt:        libraryUpdated.UpdatedAt,
	}, nil
}
