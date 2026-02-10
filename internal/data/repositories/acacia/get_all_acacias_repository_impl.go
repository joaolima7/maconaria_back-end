package acacia

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetAllAcaciasRepositoryImpl struct {
	queries *db.Queries
}

func NewGetAllAcaciasRepositoryImpl(queries *db.Queries) *GetAllAcaciasRepositoryImpl {
	return &GetAllAcaciasRepositoryImpl{queries: queries}
}

func (r *GetAllAcaciasRepositoryImpl) GetAllAcacias() ([]*entity.Acacia, error) {
	ctx := context.Background()

	acaciasDB, err := r.queries.GetAllAcacias(ctx)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar acácias")
	}

	acacias := make([]*entity.Acacia, len(acaciasDB))
	for i, acaciaDB := range acaciasDB {
		acacia, err := dbAcaciaAllRowToEntity(acaciaDB)
		if err != nil {
			return nil, err
		}
		acacias[i] = acacia
	}

	return acacias, nil
}
