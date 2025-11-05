package worker

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type DeleteWorkerRepositoryImpl struct {
	queries *db.Queries
}

func NewDeleteWorkerRepositoryImpl(queries *db.Queries) *DeleteWorkerRepositoryImpl {
	return &DeleteWorkerRepositoryImpl{queries: queries}
}

func (r *DeleteWorkerRepositoryImpl) DeleteWorker(id string) error {
	ctx := context.Background()

	_, err := r.queries.GetWorkerByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.NewNotFoundError("Obreiro")
		}
		return apperrors.WrapDatabaseError(err, "buscar obreiro!")
	}

	if err := r.queries.DeleteWorker(ctx, id); err != nil {
		return apperrors.WrapDatabaseError(err, "deletar obreiro!")
	}

	return nil
}
