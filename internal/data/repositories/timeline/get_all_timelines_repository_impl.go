package timeline

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetAllTimelinesRepositoryImpl struct {
	queries *db.Queries
}

func NewGetAllTimelinesRepositoryImpl(queries *db.Queries) *GetAllTimelinesRepositoryImpl {
	return &GetAllTimelinesRepositoryImpl{queries: queries}
}

func (r *GetAllTimelinesRepositoryImpl) GetAllTimelines() ([]*entity.Timeline, error) {
	ctx := context.Background()

	timelinesDB, err := r.queries.GetAllTimelines(ctx)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar timelines")
	}

	timelines := make([]*entity.Timeline, len(timelinesDB))
	for i, timelineDB := range timelinesDB {
		timelines[i] = dbTimelineToEntity(timelineDB)
	}

	return timelines, nil
}
