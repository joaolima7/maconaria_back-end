package timeline

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type UpdateTimelineByIDRepository interface {
	UpdateTimelineByID(timeline *entity.Timeline) (*entity.Timeline, error)
}
