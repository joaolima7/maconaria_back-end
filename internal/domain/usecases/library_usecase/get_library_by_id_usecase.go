package library_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"
)

type GetLibraryByIDInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type GetLibraryByIDOutputDTO struct {
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

type GetLibraryByIDUseCase struct {
	Repository library.GetLibraryByIDRepository
}

func NewGetLibraryByIDUseCase(repository library.GetLibraryByIDRepository) *GetLibraryByIDUseCase {
	return &GetLibraryByIDUseCase{
		Repository: repository,
	}
}

func (uc *GetLibraryByIDUseCase) Execute(input GetLibraryByIDInputDTO) (*GetLibraryByIDOutputDTO, error) {
	l, err := uc.Repository.GetLibraryByID(input.ID)
	if err != nil {
		return nil, err
	}

	fileData := ""
	if len(l.FileData) > 0 {
		fileData = base64.StdEncoding.EncodeToString(l.FileData)
	}

	coverData := ""
	if len(l.CoverData) > 0 {
		coverData = base64.StdEncoding.EncodeToString(l.CoverData)
	}

	return &GetLibraryByIDOutputDTO{
		ID:               l.ID,
		Title:            l.Title,
		SmallDescription: l.SmallDescription,
		Degree:           string(l.Degree),
		FileData:         fileData,
		CoverData:        coverData,
		Link:             l.Link,
		CreatedAt:        l.CreatedAt,
		UpdatedAt:        l.UpdatedAt,
	}, nil
}
