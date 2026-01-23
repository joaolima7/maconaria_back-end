package worker

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetAllWorkersRepositoryImpl struct {
	queries *db.Queries
}

func NewGetAllWorkersRepositoryImpl(queries *db.Queries) *GetAllWorkersRepositoryImpl {
	return &GetAllWorkersRepositoryImpl{queries: queries}
}

func (r *GetAllWorkersRepositoryImpl) GetAllWorkers() ([]*entity.Worker, error) {
	ctx := context.Background()

	workersDB, err := r.queries.GetAllWorkers(ctx)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar obreiros!")
	}

	workers := make([]*entity.Worker, len(workersDB))
	for i, workerDB := range workersDB {
		workers[i] = dbWorkerAllRowToEntity(workerDB)
	}

	return workers, nil
}
