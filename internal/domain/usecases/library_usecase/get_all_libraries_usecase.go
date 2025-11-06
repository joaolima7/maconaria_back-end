package library_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"
)

type GetAllLibrariesOutputDTO struct {
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

type GetAllLibrariesUseCase struct {
	Repository library.GetAllLibrariesRepository
}

func NewGetAllLibrariesUseCase(repository library.GetAllLibrariesRepository) *GetAllLibrariesUseCase {
	return &GetAllLibrariesUseCase{
		Repository: repository,
	}
}

func (uc *GetAllLibrariesUseCase) Execute() ([]*GetAllLibrariesOutputDTO, error) {
	libraries, err := uc.Repository.GetAllLibraries()
	if err != nil {
		return nil, err
	}

	output := make([]*GetAllLibrariesOutputDTO, len(libraries))
	for i, l := range libraries {
		fileData := ""
		if len(l.FileData) > 0 {
			fileData = base64.StdEncoding.EncodeToString(l.FileData)
		}

		coverData := ""
		if len(l.CoverData) > 0 {
			coverData = base64.StdEncoding.EncodeToString(l.CoverData)
		}

		output[i] = &GetAllLibrariesOutputDTO{
			ID:               l.ID,
			Title:            l.Title,
			SmallDescription: l.SmallDescription,
			Degree:           string(l.Degree),
			FileData:         fileData,
			CoverData:        coverData,
			Link:             l.Link,
			CreatedAt:        l.CreatedAt,
			UpdatedAt:        l.UpdatedAt,
		}
	}

	return output, nil
}
