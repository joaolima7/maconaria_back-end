package timeline

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type DeleteTimelineRepositoryImpl struct {
	queries *db.Queries
}

func NewDeleteTimelineRepositoryImpl(queries *db.Queries) *DeleteTimelineRepositoryImpl {
	return &DeleteTimelineRepositoryImpl{queries: queries}
}

func (r *DeleteTimelineRepositoryImpl) DeleteTimeline(id string) error {
	ctx := context.Background()

	_, err := r.queries.GetTimelineByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.NewNotFoundError("Timeline")
		}
		return apperrors.WrapDatabaseError(err, "buscar timeline")
	}

	if err := r.queries.DeleteTimeline(ctx, id); err != nil {
		return apperrors.WrapDatabaseError(err, "deletar timeline")
	}

	return nil
}
