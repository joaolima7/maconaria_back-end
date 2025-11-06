package library

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetLibraryByIDRepository interface {
	GetLibraryByID(id string) (*entity.Library, error)
}
