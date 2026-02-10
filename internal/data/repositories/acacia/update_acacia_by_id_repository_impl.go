package acacia

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type UpdateAcaciaByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewUpdateAcaciaByIDRepositoryImpl(queries *db.Queries) *UpdateAcaciaByIDRepositoryImpl {
	return &UpdateAcaciaByIDRepositoryImpl{queries: queries}
}

func (r *UpdateAcaciaByIDRepositoryImpl) UpdateAcaciaByID(acacia *entity.Acacia) (*entity.Acacia, error) {
	ctx := context.Background()

	_, err := r.queries.GetAcaciaByID(ctx, acacia.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Acácia")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar acácia")
	}

	termsJSON, err := acacia.TermsToJSON()
	if err != nil {
		return nil, apperrors.NewValidationError("mandatos", "Formato de mandatos inválido!")
	}

	params := db.UpdateAcaciaParams{
		Name:        acacia.Name,
		Terms:       []byte(termsJSON),
		IsPresident: acacia.IsPresident,
		Deceased:    acacia.Deceased,
		ImageUrl:    acacia.ImageURL,
		IsActive:    acacia.IsActive,
		ID:          acacia.ID,
	}

	_, err = r.queries.UpdateAcacia(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "atualizar acácia")
	}

	acaciaDB, err := r.queries.GetAcaciaByID(ctx, acacia.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar acácia atualizada")
	}

	return dbAcaciaByIDRowToEntity(acaciaDB)
}
