package timeline_usecase

import "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"

type DeleteTimelineInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeleteTimelineUseCase struct {
	Repository timeline.DeleteTimelineRepository
}

func NewDeleteTimelineUseCase(repository timeline.DeleteTimelineRepository) *DeleteTimelineUseCase {
	return &DeleteTimelineUseCase{
		Repository: repository,
	}
}

func (uc *DeleteTimelineUseCase) Execute(input DeleteTimelineInputDTO) error {
	return uc.Repository.DeleteTimeline(input.ID)
}
