package library_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/library"
)

type GetAllLibrariesOutputDTO struct {
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
		output[i] = &GetAllLibrariesOutputDTO{
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
