package library_usecase

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type CreateLibraryInputDTO struct {
	Title            string `json:"title" validate:"required"`
	SmallDescription string `json:"small_description" validate:"required"`
	Degree           string `json:"degree" validate:"required,oneof=apprentice companion master"`
	FileData         string `json:"file_data,omitempty" validate:"omitempty,base64"`
	CoverData        string `json:"cover_data,omitempty" validate:"omitempty,base64"`
	Link             string `json:"link,omitempty"`
}

type CreateLibraryOutputDTO struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	SmallDescription string    `json:"small_description"`
	Degree           string    `json:"degree"`
	FileURL          string    `json:"file_url,omitempty"`
	CoverURL         string    `json:"cover_url,omitempty"`
	Link             string    `json:"link,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreateLibraryUseCase struct {
	Repository     library.CreateLibraryRepository
	StorageService storage.StorageService
}

func NewCreateLibraryUseCase(
	repository library.CreateLibraryRepository,
	storageService storage.StorageService,
) *CreateLibraryUseCase {
	return &CreateLibraryUseCase{
		Repository:     repository,
		StorageService: storageService,
	}
}

func (uc *CreateLibraryUseCase) Execute(input CreateLibraryInputDTO) (*CreateLibraryOutputDTO, error) {

	if input.FileData == "" && input.Link == "" {
		return nil, apperrors.NewValidationError("conteúdo", "É necessário fornecer um arquivo PDF ou um link!")
	}

	libraryID := uuid.New().String()
	var fileURL, coverURL string
	var uploadedFiles []string

	if input.FileData != "" {
		fileData, err := base64.StdEncoding.DecodeString(input.FileData)
		if err != nil {
			return nil, apperrors.NewValidationError("arquivo", "Arquivo em formato inválido!")
		}

		filename := fmt.Sprintf("library_%s_%s.pdf", libraryID, uuid.New().String())
		fileURL, err = uc.StorageService.UploadPDF(fileData, filename, "libraries")
		if err != nil {
			return nil, apperrors.NewInternalError("Erro ao fazer upload do arquivo", err)
		}
		uploadedFiles = append(uploadedFiles, fileURL)
	}

	if input.CoverData != "" {
		coverData, err := base64.StdEncoding.DecodeString(input.CoverData)
		if err != nil {

			for _, url := range uploadedFiles {
				_ = uc.StorageService.DeletePDF(url, "libraries")
			}
			return nil, apperrors.NewValidationError("capa", "Capa em formato inválido!")
		}

		filename := fmt.Sprintf("library_cover_%s_%s.jpg", libraryID, uuid.New().String())
		coverURL, err = uc.StorageService.UploadImage(coverData, filename, "libraries")
		if err != nil {

			for _, url := range uploadedFiles {
				_ = uc.StorageService.DeletePDF(url, "libraries")
			}
			return nil, apperrors.NewInternalError("Erro ao fazer upload da capa", err)
		}
		uploadedFiles = append(uploadedFiles, coverURL)
	}

	libraryEntity, err := entity.NewLibrary(
		libraryID,
		input.Title,
		input.SmallDescription,
		entity.UserDegree(input.Degree),
		fileURL,
		coverURL,
		input.Link,
	)
	if err != nil {

		for _, url := range uploadedFiles {
			if url == fileURL {
				_ = uc.StorageService.DeletePDF(url, "libraries")
			} else {
				_ = uc.StorageService.DeleteImage(url, "libraries")
			}
		}
		return nil, err
	}

	libraryCreated, err := uc.Repository.CreateLibrary(libraryEntity)
	if err != nil {

		for _, url := range uploadedFiles {
			if url == fileURL {
				_ = uc.StorageService.DeletePDF(url, "libraries")
			} else {
				_ = uc.StorageService.DeleteImage(url, "libraries")
			}
		}
		return nil, err
	}

	return &CreateLibraryOutputDTO{
		ID:               libraryCreated.ID,
		Title:            libraryCreated.Title,
		SmallDescription: libraryCreated.SmallDescription,
		Degree:           string(libraryCreated.Degree),
		FileURL:          libraryCreated.FileURL,
		CoverURL:         libraryCreated.CoverURL,
		Link:             libraryCreated.Link,
		CreatedAt:        libraryCreated.CreatedAt,
		UpdatedAt:        libraryCreated.UpdatedAt,
	}, nil
}
