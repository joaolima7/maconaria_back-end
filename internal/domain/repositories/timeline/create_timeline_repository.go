package timeline

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type CreateTimelineRepository interface {
	CreateTimeline(timeline *entity.Timeline) (*entity.Timeline, error)
}
