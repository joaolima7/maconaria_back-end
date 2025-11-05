package timeline

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetTimelineByIDRepository interface {
	GetTimelineByID(id string) (*entity.Timeline, error)
}
