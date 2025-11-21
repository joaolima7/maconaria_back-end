package timeline

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type CreateTimelineRepositoryImpl struct {
	queries *db.Queries
}

func NewCreateTimelineRepositoryImpl(queries *db.Queries) *CreateTimelineRepositoryImpl {
	return &CreateTimelineRepositoryImpl{queries: queries}
}

func (r *CreateTimelineRepositoryImpl) CreateTimeline(timeline *entity.Timeline) (*entity.Timeline, error) {
	ctx := context.Background()

	existingByPeriod, err := r.queries.GetTimelineByPeriod(ctx, timeline.Period)
	if err == nil && existingByPeriod.Period == timeline.Period {
		return nil, apperrors.NewDuplicateError("período", timeline.Period)
	}

	params := db.CreateTimelineParams{
		ID:          timeline.ID,
		Period:      timeline.Period,
		PdfUrl:      timeline.PdfURL,
		IsHighlight: timeline.IsHighlight,
	}

	_, err = r.queries.CreateTimeline(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "criar timeline")
	}

	timelineDB, err := r.queries.GetTimelineByID(ctx, timeline.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar timeline criada")
	}

	return dbTimelineToEntity(timelineDB), nil
}

func dbTimelineToEntity(timelineDB db.Timeline) *entity.Timeline {
	return &entity.Timeline{
		ID:          timelineDB.ID,
		Period:      timelineDB.Period,
		PdfURL:      timelineDB.PdfUrl,
		IsHighlight: timelineDB.IsHighlight,
		CreatedAt:   timelineDB.CreatedAt.Time,
		UpdatedAt:   timelineDB.UpdatedAt.Time,
	}
}
