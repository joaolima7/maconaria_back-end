package library_usecase

import (
	"encoding/base64"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"
)

type GetLibrariesByDegreeInputDTO struct {
	Degree string `json:"degree" validate:"required,oneof=apprentice companion master"`
}

type GetLibrariesByDegreeOutputDTO struct {
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

type GetLibrariesByDegreeUseCase struct {
	Repository library.GetLibrariesByDegreeRepository
}

func NewGetLibrariesByDegreeUseCase(repository library.GetLibrariesByDegreeRepository) *GetLibrariesByDegreeUseCase {
	return &GetLibrariesByDegreeUseCase{
		Repository: repository,
	}
}

func (uc *GetLibrariesByDegreeUseCase) Execute(input GetLibrariesByDegreeInputDTO) ([]*GetLibrariesByDegreeOutputDTO, error) {
	libraries, err := uc.Repository.GetLibrariesByDegree(input.Degree)
	if err != nil {
		return nil, err
	}

	output := make([]*GetLibrariesByDegreeOutputDTO, len(libraries))
	for i, l := range libraries {
		fileData := ""
		if len(l.FileData) > 0 {
			fileData = base64.StdEncoding.EncodeToString(l.FileData)
		}

		coverData := ""
		if len(l.CoverData) > 0 {
			coverData = base64.StdEncoding.EncodeToString(l.CoverData)
		}

		output[i] = &GetLibrariesByDegreeOutputDTO{
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
