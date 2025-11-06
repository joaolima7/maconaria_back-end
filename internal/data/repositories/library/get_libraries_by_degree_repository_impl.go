package library

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetLibrariesByDegreeRepositoryImpl struct {
	queries *db.Queries
}

func NewGetLibrariesByDegreeRepositoryImpl(queries *db.Queries) *GetLibrariesByDegreeRepositoryImpl {
	return &GetLibrariesByDegreeRepositoryImpl{queries: queries}
}

func (r *GetLibrariesByDegreeRepositoryImpl) GetLibrariesByDegree(degree string) ([]*entity.Library, error) {
	ctx := context.Background()

	librariesDB, err := r.queries.GetLibrariesByDegree(ctx, db.LibrariesDegree(degree))
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar bibliotecas por grau")
	}

	libraries := make([]*entity.Library, len(librariesDB))
	for i, libraryDB := range librariesDB {
		libraries[i] = dbLibraryToEntity(libraryDB)
	}

	return libraries, nil
}
