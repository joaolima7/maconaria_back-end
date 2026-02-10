package acacia

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetAcaciaByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewGetAcaciaByIDRepositoryImpl(queries *db.Queries) *GetAcaciaByIDRepositoryImpl {
	return &GetAcaciaByIDRepositoryImpl{queries: queries}
}

func (r *GetAcaciaByIDRepositoryImpl) GetAcaciaByID(id string) (*entity.Acacia, error) {
	ctx := context.Background()

	acaciaDB, err := r.queries.GetAcaciaByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Acácia")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar acácia")
	}

	return dbAcaciaByIDRowToEntity(acaciaDB)
}
