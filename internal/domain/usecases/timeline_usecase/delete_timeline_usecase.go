package timeline_usecase

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type DeleteTimelineInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeleteTimelineUseCase struct {
	Repository     timeline.DeleteTimelineRepository
	GetRepository  timeline.GetTimelineByIDRepository
	StorageService storage.StorageService
}

func NewDeleteTimelineUseCase(
	repository timeline.DeleteTimelineRepository,
	getRepository timeline.GetTimelineByIDRepository,
	storageService storage.StorageService,
) *DeleteTimelineUseCase {
	return &DeleteTimelineUseCase{
		Repository:     repository,
		GetRepository:  getRepository,
		StorageService: storageService,
	}
}

func (uc *DeleteTimelineUseCase) Execute(input DeleteTimelineInputDTO) error {

	t, err := uc.GetRepository.GetTimelineByID(input.ID)
	if err != nil {
		return err
	}

	if err := uc.Repository.DeleteTimeline(input.ID); err != nil {
		return err
	}

	_ = uc.StorageService.DeletePDF(t.PdfURL, "timelines")

	return nil
}
