package worker

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type CreateWorkerRepositoryImpl struct {
	queries *db.Queries
}

func NewCreateWorkerRepositoryImpl(queries *db.Queries) *CreateWorkerRepositoryImpl {
	return &CreateWorkerRepositoryImpl{queries: queries}
}

func (r *CreateWorkerRepositoryImpl) CreateWorker(worker *entity.Worker) (*entity.Worker, error) {
	ctx := context.Background()

	existingByNumber, err := r.queries.GetWorkerByNumber(ctx, worker.Number)
	if err == nil && existingByNumber.Number == worker.Number {
		return nil, apperrors.NewDuplicateError("número", fmt.Sprintf("%d", worker.Number))
	}

	existingByRegistration, err := r.queries.GetWorkerByRegistration(ctx, worker.Registration)
	if err == nil && existingByRegistration.Registration == worker.Registration {
		return nil, apperrors.NewDuplicateError("cadastro", worker.Registration)
	}

	var emeritusMasonDate, provectMasonDate sql.NullTime
	if worker.EmeritusMasonDate != nil {
		emeritusMasonDate = sql.NullTime{Time: *worker.EmeritusMasonDate, Valid: true}
	}
	if worker.ProvectMasonDate != nil {
		provectMasonDate = sql.NullTime{Time: *worker.ProvectMasonDate, Valid: true}
	}

	params := db.CreateWorkerParams{
		ID:                worker.ID,
		Number:            worker.Number,
		Name:              worker.Name,
		Registration:      worker.Registration,
		BirthDate:         worker.BirthDate,
		InitiationDate:    worker.InitiationDate,
		ElevationDate:     worker.ElevationDate,
		ExaltationDate:    worker.ExaltationDate,
		AffiliationDate:   worker.AffiliationDate,
		InstallationDate:  worker.InstallationDate,
		EmeritusMasonDate: emeritusMasonDate,
		ProvectMasonDate:  provectMasonDate,
		ImageData:         worker.ImageData,
		Deceased:          worker.Deceased,
	}

	_, err = r.queries.CreateWorker(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "criar obreiro")
	}

	workerDB, err := r.queries.GetWorkerByID(ctx, worker.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar obreiro criado")
	}

	return dbWorkerToEntity(workerDB), nil
}

func dbWorkerToEntity(workerDB db.Worker) *entity.Worker {
	var emeritusMasonDate, provectMasonDate *time.Time
	if workerDB.EmeritusMasonDate.Valid {
		emeritusMasonDate = &workerDB.EmeritusMasonDate.Time
	}
	if workerDB.ProvectMasonDate.Valid {
		provectMasonDate = &workerDB.ProvectMasonDate.Time
	}

	return &entity.Worker{
		ID:                workerDB.ID,
		Number:            workerDB.Number,
		Name:              workerDB.Name,
		Registration:      workerDB.Registration,
		BirthDate:         workerDB.BirthDate,
		InitiationDate:    workerDB.InitiationDate,
		ElevationDate:     workerDB.ElevationDate,
		ExaltationDate:    workerDB.ExaltationDate,
		AffiliationDate:   workerDB.AffiliationDate,
		InstallationDate:  workerDB.InstallationDate,
		EmeritusMasonDate: emeritusMasonDate,
		ProvectMasonDate:  provectMasonDate,
		ImageData:         workerDB.ImageData,
		Deceased:          workerDB.Deceased,
		CreatedAt:         workerDB.CreatedAt.Time,
		UpdatedAt:         workerDB.UpdatedAt.Time,
	}
}
