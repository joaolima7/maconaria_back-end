package timeline

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetTimelineByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewGetTimelineByIDRepositoryImpl(queries *db.Queries) *GetTimelineByIDRepositoryImpl {
	return &GetTimelineByIDRepositoryImpl{queries: queries}
}

func (r *GetTimelineByIDRepositoryImpl) GetTimelineByID(id string) (*entity.Timeline, error) {
	ctx := context.Background()

	timelineDB, err := r.queries.GetTimelineByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Timeline")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar timeline")
	}

	return dbTimelineToEntity(timelineDB), nil
}
