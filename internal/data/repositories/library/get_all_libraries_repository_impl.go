package library

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetAllLibrariesRepositoryImpl struct {
	queries *db.Queries
}

func NewGetAllLibrariesRepositoryImpl(queries *db.Queries) *GetAllLibrariesRepositoryImpl {
	return &GetAllLibrariesRepositoryImpl{queries: queries}
}

func (r *GetAllLibrariesRepositoryImpl) GetAllLibraries() ([]*entity.Library, error) {
	ctx := context.Background()

	librariesDB, err := r.queries.GetAllLibraries(ctx)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar bibliotecas")
	}

	libraries := make([]*entity.Library, len(librariesDB))
	for i, libraryDB := range librariesDB {
		libraries[i] = dbLibraryToEntity(libraryDB)
	}

	return libraries, nil
}
