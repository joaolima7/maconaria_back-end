package library

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetLibraryByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewGetLibraryByIDRepositoryImpl(queries *db.Queries) *GetLibraryByIDRepositoryImpl {
	return &GetLibraryByIDRepositoryImpl{queries: queries}
}

func (r *GetLibraryByIDRepositoryImpl) GetLibraryByID(id string) (*entity.Library, error) {
	ctx := context.Background()

	libraryDB, err := r.queries.GetLibraryByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Biblioteca")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar biblioteca")
	}

	return dbLibraryToEntity(libraryDB), nil
}
