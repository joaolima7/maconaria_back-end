package post_usecase

import (
	post_repository "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/post"
	"github.com/joaolima7/maconaria_back-end/internal/infra/storage"
)

type DeletePostInputDTO struct {
	ID string `json:"id" validate:"required"`
}

type DeletePostUseCase struct {
	Repository      post_repository.DeletePostRepository
	ImageRepository post_repository.PostImageRepository
	StorageService  storage.StorageService
}

func NewDeletePostUseCase(
	repository post_repository.DeletePostRepository,
	imageRepository post_repository.PostImageRepository,
	storageService storage.StorageService,
) *DeletePostUseCase {
	return &DeletePostUseCase{
		Repository:      repository,
		ImageRepository: imageRepository,
		StorageService:  storageService,
	}
}

func (uc *DeletePostUseCase) Execute(input DeletePostInputDTO) error {

	images, err := uc.ImageRepository.GetPostImages(input.ID)
	if err != nil {
		return err
	}

	if err := uc.Repository.Delete(input.ID); err != nil {
		return err
	}

	for _, img := range images {
		_ = uc.StorageService.DeleteImage(img.ImageURL, "posts")
	}

	return nil
}
