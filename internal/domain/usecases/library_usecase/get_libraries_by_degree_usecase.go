package library_usecase

import (
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
	FileURL          string    `json:"file_url,omitempty"`
	CoverURL         string    `json:"cover_url,omitempty"`
	Link             string    `json:"link,omitempty"`
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
		output[i] = &GetLibrariesByDegreeOutputDTO{
			ID:               l.ID,
			Title:            l.Title,
			SmallDescription: l.SmallDescription,
			Degree:           string(l.Degree),
			FileURL:          l.FileURL,
			CoverURL:         l.CoverURL,
			Link:             l.Link,
			CreatedAt:        l.CreatedAt,
			UpdatedAt:        l.UpdatedAt,
		}
	}

	return output, nil
}
