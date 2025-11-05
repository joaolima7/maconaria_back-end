package acacia

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type DeleteAcaciaRepositoryImpl struct {
	queries *db.Queries
}

func NewDeleteAcaciaRepositoryImpl(queries *db.Queries) *DeleteAcaciaRepositoryImpl {
	return &DeleteAcaciaRepositoryImpl{queries: queries}
}

func (r *DeleteAcaciaRepositoryImpl) DeleteAcacia(id string) error {
	ctx := context.Background()

	_, err := r.queries.GetAcaciaByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.NewNotFoundError("Acácia")
		}
		return apperrors.WrapDatabaseError(err, "buscar acácia")
	}

	if err := r.queries.DeleteAcacia(ctx, id); err != nil {
		return apperrors.WrapDatabaseError(err, "deletar acácia")
	}

	return nil
}
