package library

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetLibrariesByDegreeRepository interface {
	GetLibrariesByDegree(degree string) ([]*entity.Library, error)
}
