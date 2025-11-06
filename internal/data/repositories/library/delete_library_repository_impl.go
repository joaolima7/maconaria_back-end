package library

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type DeleteLibraryRepositoryImpl struct {
	queries *db.Queries
}

func NewDeleteLibraryRepositoryImpl(queries *db.Queries) *DeleteLibraryRepositoryImpl {
	return &DeleteLibraryRepositoryImpl{queries: queries}
}

func (r *DeleteLibraryRepositoryImpl) DeleteLibrary(id string) error {
	ctx := context.Background()

	_, err := r.queries.GetLibraryByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.NewNotFoundError("Biblioteca")
		}
		return apperrors.WrapDatabaseError(err, "buscar biblioteca")
	}

	if err := r.queries.DeleteLibrary(ctx, id); err != nil {
		return apperrors.WrapDatabaseError(err, "deletar biblioteca")
	}

	return nil
}
