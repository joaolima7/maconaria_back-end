package timeline

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetAllTimelinesRepository interface {
	GetAllTimelines() ([]*entity.Timeline, error)
}
