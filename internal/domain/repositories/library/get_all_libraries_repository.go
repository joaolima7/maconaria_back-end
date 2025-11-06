package library

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetAllLibrariesRepository interface {
	GetAllLibraries() ([]*entity.Library, error)
}
