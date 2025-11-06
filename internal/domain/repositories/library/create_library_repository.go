package library

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type CreateLibraryRepository interface {
	CreateLibrary(library *entity.Library) (*entity.Library, error)
}
