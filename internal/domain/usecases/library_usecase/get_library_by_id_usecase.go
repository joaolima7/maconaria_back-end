package library_usecase

import (
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
	FileURL          string    `json:"file_url,omitempty"`
	CoverURL         string    `json:"cover_url,omitempty"`
	Link             string    `json:"link,omitempty"`
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

	return &GetLibraryByIDOutputDTO{
		ID:               l.ID,
		Title:            l.Title,
		SmallDescription: l.SmallDescription,
		Degree:           string(l.Degree),
		FileURL:          l.FileURL,
		CoverURL:         l.CoverURL,
		Link:             l.Link,
		CreatedAt:        l.CreatedAt,
		UpdatedAt:        l.UpdatedAt,
	}, nil
}
