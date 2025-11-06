package library

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type UpdateLibraryByIDRepository interface {
	UpdateLibraryByID(library *entity.Library) (*entity.Library, error)
}
