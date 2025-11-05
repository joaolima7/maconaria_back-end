package timeline

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type UpdateTimelineByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewUpdateTimelineByIDRepositoryImpl(queries *db.Queries) *UpdateTimelineByIDRepositoryImpl {
	return &UpdateTimelineByIDRepositoryImpl{queries: queries}
}

func (r *UpdateTimelineByIDRepositoryImpl) UpdateTimelineByID(timeline *entity.Timeline) (*entity.Timeline, error) {
	ctx := context.Background()

	_, err := r.queries.GetTimelineByID(ctx, timeline.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Timeline")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar timeline")
	}

	params := db.UpdateTimelineParams{
		Period:      timeline.Period,
		PdfData:     timeline.PdfData,
		IsHighlight: timeline.IsHighlight,
		ID:          timeline.ID,
	}

	_, err = r.queries.UpdateTimeline(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "atualizar timeline")
	}

	timelineDB, err := r.queries.GetTimelineByID(ctx, timeline.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar timeline atualizada")
	}

	return dbTimelineToEntity(timelineDB), nil
}
