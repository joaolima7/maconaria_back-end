package worker

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type UpdateWorkerByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewUpdateWorkerByIDRepositoryImpl(queries *db.Queries) *UpdateWorkerByIDRepositoryImpl {
	return &UpdateWorkerByIDRepositoryImpl{queries: queries}
}

func (r *UpdateWorkerByIDRepositoryImpl) UpdateWorkerByID(worker *entity.Worker) (*entity.Worker, error) {
	ctx := context.Background()

	_, err := r.queries.GetWorkerByID(ctx, worker.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Obreiro")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar obreiro!")
	}

	var emeritusMasonDate, provectMasonDate sql.NullTime
	if worker.EmeritusMasonDate != nil {
		emeritusMasonDate = sql.NullTime{Time: *worker.EmeritusMasonDate, Valid: true}
	}
	if worker.ProvectMasonDate != nil {
		provectMasonDate = sql.NullTime{Time: *worker.ProvectMasonDate, Valid: true}
	}

	var initiationDate, elevationDate, exaltationDate, affiliationDate, installationDate sql.NullTime
	if worker.InitiationDate != nil {
		initiationDate = sql.NullTime{Time: *worker.InitiationDate, Valid: true}
	}
	if worker.ElevationDate != nil {
		elevationDate = sql.NullTime{Time: *worker.ElevationDate, Valid: true}
	}
	if worker.ExaltationDate != nil {
		exaltationDate = sql.NullTime{Time: *worker.ExaltationDate, Valid: true}
	}
	if worker.AffiliationDate != nil {
		affiliationDate = sql.NullTime{Time: *worker.AffiliationDate, Valid: true}
	}
	if worker.InstallationDate != nil {
		installationDate = sql.NullTime{Time: *worker.InstallationDate, Valid: true}
	}

	params := db.UpdateWorkerParams{
		Number:            worker.Number,
		Name:              worker.Name,
		Registration:      worker.Registration,
		BirthDate:         worker.BirthDate,
		InitiationDate:    initiationDate,
		ElevationDate:     elevationDate,
		ExaltationDate:    exaltationDate,
		AffiliationDate:   affiliationDate,
		InstallationDate:  installationDate,
		EmeritusMasonDate: emeritusMasonDate,
		ProvectMasonDate:  provectMasonDate,
		ImageUrl:          worker.ImageURL,
		Deceased:          worker.Deceased,
		ID:                worker.ID,
	}

	_, err = r.queries.UpdateWorker(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "atualizar obreiro!")
	}

	workerDB, err := r.queries.GetWorkerByID(ctx, worker.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar obreiro atualizado!")
	}

	return dbWorkerToEntity(workerDB), nil
}
