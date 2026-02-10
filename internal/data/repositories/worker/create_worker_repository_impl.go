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

	termsJSON, err := worker.TermsToJSON()
	if err != nil {
		return nil, apperrors.NewValidationError("mandatos", "Erro ao processar mandatos!")
	}

	params := db.CreateWorkerParams{
		ID:                worker.ID,
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
		IsPresident:       worker.IsPresident,
		Terms:             []byte(termsJSON),
		IsActive:          worker.IsActive,
	}

	_, err = r.queries.CreateWorker(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "criar obreiro")
	}

	workerDB, err := r.queries.GetWorkerByID(ctx, worker.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar obreiro criado")
	}

	return dbWorkerRowToEntity(workerDB), nil
}

func dbWorkerToEntity(workerDB db.Worker) *entity.Worker {
	var initiationDate, elevationDate, exaltationDate, affiliationDate, installationDate *time.Time
	var emeritusMasonDate, provectMasonDate *time.Time

	if workerDB.InitiationDate.Valid {
		initiationDate = &workerDB.InitiationDate.Time
	}
	if workerDB.ElevationDate.Valid {
		elevationDate = &workerDB.ElevationDate.Time
	}
	if workerDB.ExaltationDate.Valid {
		exaltationDate = &workerDB.ExaltationDate.Time
	}
	if workerDB.AffiliationDate.Valid {
		affiliationDate = &workerDB.AffiliationDate.Time
	}
	if workerDB.InstallationDate.Valid {
		installationDate = &workerDB.InstallationDate.Time
	}
	if workerDB.EmeritusMasonDate.Valid {
		emeritusMasonDate = &workerDB.EmeritusMasonDate.Time
	}
	if workerDB.ProvectMasonDate.Valid {
		provectMasonDate = &workerDB.ProvectMasonDate.Time
	}

	terms, err := entity.TermsFromJSON(string(workerDB.Terms))
	if err != nil {
		terms = []string{}
	}

	return &entity.Worker{
		ID:                workerDB.ID,
		Number:            workerDB.Number,
		Name:              workerDB.Name,
		Registration:      workerDB.Registration,
		BirthDate:         workerDB.BirthDate,
		InitiationDate:    initiationDate,
		ElevationDate:     elevationDate,
		ExaltationDate:    exaltationDate,
		AffiliationDate:   affiliationDate,
		InstallationDate:  installationDate,
		EmeritusMasonDate: emeritusMasonDate,
		ProvectMasonDate:  provectMasonDate,
		ImageURL:          workerDB.ImageUrl,
		Deceased:          workerDB.Deceased,
		IsPresident:       workerDB.IsPresident,
		Terms:             terms,
		IsActive:          workerDB.IsActive,
		CreatedAt:         workerDB.CreatedAt.Time,
		UpdatedAt:         workerDB.UpdatedAt.Time,
	}
}

func dbWorkerRowToEntity(workerDB db.GetWorkerByIDRow) *entity.Worker {
	var initiationDate, elevationDate, exaltationDate, affiliationDate, installationDate *time.Time
	var emeritusMasonDate, provectMasonDate *time.Time

	if workerDB.InitiationDate.Valid {
		initiationDate = &workerDB.InitiationDate.Time
	}
	if workerDB.ElevationDate.Valid {
		elevationDate = &workerDB.ElevationDate.Time
	}
	if workerDB.ExaltationDate.Valid {
		exaltationDate = &workerDB.ExaltationDate.Time
	}
	if workerDB.AffiliationDate.Valid {
		affiliationDate = &workerDB.AffiliationDate.Time
	}
	if workerDB.InstallationDate.Valid {
		installationDate = &workerDB.InstallationDate.Time
	}
	if workerDB.EmeritusMasonDate.Valid {
		emeritusMasonDate = &workerDB.EmeritusMasonDate.Time
	}
	if workerDB.ProvectMasonDate.Valid {
		provectMasonDate = &workerDB.ProvectMasonDate.Time
	}

	terms, err := entity.TermsFromJSON(string(workerDB.Terms))
	if err != nil {
		terms = []string{}
	}

	return &entity.Worker{
		ID:                workerDB.ID,
		Number:            workerDB.Number,
		Name:              workerDB.Name,
		Registration:      workerDB.Registration,
		BirthDate:         workerDB.BirthDate,
		InitiationDate:    initiationDate,
		ElevationDate:     elevationDate,
		ExaltationDate:    exaltationDate,
		AffiliationDate:   affiliationDate,
		InstallationDate:  installationDate,
		EmeritusMasonDate: emeritusMasonDate,
		ProvectMasonDate:  provectMasonDate,
		ImageURL:          workerDB.ImageUrl,
		Deceased:          workerDB.Deceased,
		IsPresident:       workerDB.IsPresident,
		Terms:             terms,
		IsActive:          workerDB.IsActive,
		CreatedAt:         workerDB.CreatedAt.Time,
		UpdatedAt:         workerDB.UpdatedAt.Time,
	}
}

func dbWorkerAllRowToEntity(workerDB db.GetAllWorkersRow) *entity.Worker {
	var initiationDate, elevationDate, exaltationDate, affiliationDate, installationDate *time.Time
	var emeritusMasonDate, provectMasonDate *time.Time

	if workerDB.InitiationDate.Valid {
		initiationDate = &workerDB.InitiationDate.Time
	}
	if workerDB.ElevationDate.Valid {
		elevationDate = &workerDB.ElevationDate.Time
	}
	if workerDB.ExaltationDate.Valid {
		exaltationDate = &workerDB.ExaltationDate.Time
	}
	if workerDB.AffiliationDate.Valid {
		affiliationDate = &workerDB.AffiliationDate.Time
	}
	if workerDB.InstallationDate.Valid {
		installationDate = &workerDB.InstallationDate.Time
	}
	if workerDB.EmeritusMasonDate.Valid {
		emeritusMasonDate = &workerDB.EmeritusMasonDate.Time
	}
	if workerDB.ProvectMasonDate.Valid {
		provectMasonDate = &workerDB.ProvectMasonDate.Time
	}

	terms, err := entity.TermsFromJSON(string(workerDB.Terms))
	if err != nil {
		terms = []string{}
	}

	return &entity.Worker{
		ID:                workerDB.ID,
		Number:            workerDB.Number,
		Name:              workerDB.Name,
		Registration:      workerDB.Registration,
		BirthDate:         workerDB.BirthDate,
		InitiationDate:    initiationDate,
		ElevationDate:     elevationDate,
		ExaltationDate:    exaltationDate,
		AffiliationDate:   affiliationDate,
		InstallationDate:  installationDate,
		EmeritusMasonDate: emeritusMasonDate,
		ProvectMasonDate:  provectMasonDate,
		ImageURL:          workerDB.ImageUrl,
		Deceased:          workerDB.Deceased,
		IsPresident:       workerDB.IsPresident,
		Terms:             terms,
		CreatedAt:         workerDB.CreatedAt.Time,
		UpdatedAt:         workerDB.UpdatedAt.Time,
	}
}
