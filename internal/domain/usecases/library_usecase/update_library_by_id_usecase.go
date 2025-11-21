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

type UpdateLibraryByIDInputDTO struct {
	ID               string `json:"id"`
	Title            string `json:"title" validate:"required"`
	SmallDescription string `json:"small_description" validate:"required"`
	Degree           string `json:"degree" validate:"required,oneof=apprentice companion master"`
	FileData         string `json:"file_data,omitempty"`
	CoverData        string `json:"cover_data,omitempty"`
	Link             string `json:"link,omitempty"`
}

type UpdateLibraryByIDOutputDTO struct {
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

type UpdateLibraryByIDUseCase struct {
	Repository     library.UpdateLibraryByIDRepository
	GetRepository  library.GetLibraryByIDRepository
	StorageService storage.StorageService
}

func NewUpdateLibraryByIDUseCase(
	repository library.UpdateLibraryByIDRepository,
	getRepository library.GetLibraryByIDRepository,
	storageService storage.StorageService,
) *UpdateLibraryByIDUseCase {
	return &UpdateLibraryByIDUseCase{
		Repository:     repository,
		GetRepository:  getRepository,
		StorageService: storageService,
	}
}

func (uc *UpdateLibraryByIDUseCase) Execute(input UpdateLibraryByIDInputDTO) (*UpdateLibraryByIDOutputDTO, error) {

	existingLibrary, err := uc.GetRepository.GetLibraryByID(input.ID)
	if err != nil {
		return nil, err
	}

	fileURL := existingLibrary.FileURL
	coverURL := existingLibrary.CoverURL
	var oldFileURL, oldCoverURL string
	var uploadedFiles []string

	if input.FileData != "" {
		fileData, err := base64.StdEncoding.DecodeString(input.FileData)
		if err != nil {
			return nil, apperrors.NewValidationError("arquivo", "Arquivo em formato inválido!")
		}

		filename := fmt.Sprintf("library_%s_%s.pdf", input.ID, uuid.New().String())
		newFileURL, err := uc.StorageService.UploadPDF(fileData, filename, "libraries")
		if err != nil {
			return nil, apperrors.NewInternalError("Erro ao fazer upload do arquivo", err)
		}

		oldFileURL = fileURL
		fileURL = newFileURL
		uploadedFiles = append(uploadedFiles, newFileURL)
	}

	if input.CoverData != "" {
		coverData, err := base64.StdEncoding.DecodeString(input.CoverData)
		if err != nil {

			for _, url := range uploadedFiles {
				_ = uc.StorageService.DeletePDF(url, "libraries")
			}
			return nil, apperrors.NewValidationError("capa", "Capa em formato inválido!")
		}

		filename := fmt.Sprintf("library_cover_%s_%s.jpg", input.ID, uuid.New().String())
		newCoverURL, err := uc.StorageService.UploadImage(coverData, filename, "libraries")
		if err != nil {

			for _, url := range uploadedFiles {
				_ = uc.StorageService.DeletePDF(url, "libraries")
			}
			return nil, apperrors.NewInternalError("Erro ao fazer upload da capa", err)
		}

		oldCoverURL = coverURL
		coverURL = newCoverURL
		uploadedFiles = append(uploadedFiles, newCoverURL)
	}

	libraryEntity := &entity.Library{
		ID:               input.ID,
		Title:            input.Title,
		SmallDescription: input.SmallDescription,
		Degree:           entity.UserDegree(input.Degree),
		FileURL:          fileURL,
		CoverURL:         coverURL,
		Link:             input.Link,
		UpdatedAt:        time.Now(),
	}

	if err := libraryEntity.Validate(); err != nil {

		for _, url := range uploadedFiles {
			if url == fileURL && url != oldFileURL {
				_ = uc.StorageService.DeletePDF(url, "libraries")
			} else if url == coverURL && url != oldCoverURL {
				_ = uc.StorageService.DeleteImage(url, "libraries")
			}
		}
		return nil, err
	}

	libraryUpdated, err := uc.Repository.UpdateLibraryByID(libraryEntity)
	if err != nil {

		for _, url := range uploadedFiles {
			if url == fileURL && url != oldFileURL {
				_ = uc.StorageService.DeletePDF(url, "libraries")
			} else if url == coverURL && url != oldCoverURL {
				_ = uc.StorageService.DeleteImage(url, "libraries")
			}
		}
		return nil, err
	}

	if oldFileURL != "" {
		_ = uc.StorageService.DeletePDF(oldFileURL, "libraries")
	}
	if oldCoverURL != "" {
		_ = uc.StorageService.DeleteImage(oldCoverURL, "libraries")
	}

	return &UpdateLibraryByIDOutputDTO{
		ID:               libraryUpdated.ID,
		Title:            libraryUpdated.Title,
		SmallDescription: libraryUpdated.SmallDescription,
		Degree:           string(libraryUpdated.Degree),
		FileURL:          libraryUpdated.FileURL,
		CoverURL:         libraryUpdated.CoverURL,
		Link:             libraryUpdated.Link,
		CreatedAt:        libraryUpdated.CreatedAt,
		UpdatedAt:        libraryUpdated.UpdatedAt,
	}, nil
}
