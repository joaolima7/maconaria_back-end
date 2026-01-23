package worker

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetWorkerByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewGetWorkerByIDRepositoryImpl(queries *db.Queries) *GetWorkerByIDRepositoryImpl {
	return &GetWorkerByIDRepositoryImpl{queries: queries}
}

func (r *GetWorkerByIDRepositoryImpl) GetWorkerByID(id string) (*entity.Worker, error) {
	ctx := context.Background()

	workerDB, err := r.queries.GetWorkerByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Obreiro")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar obreiro!")
	}

	return dbWorkerRowToEntity(workerDB), nil
}
