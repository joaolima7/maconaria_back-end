package post

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type PostImageRepositoryImpl struct {
	queries *db.Queries
}

func NewPostImageRepositoryImpl(queries *db.Queries) *PostImageRepositoryImpl {
	return &PostImageRepositoryImpl{queries: queries}
}

func (r *PostImageRepositoryImpl) CreatePostImage(image *entity.PostImage) error {
	ctx := context.Background()

	params := db.CreatePostImageParams{
		ID:        image.ID,
		PostID:    image.PostID,
		ImageData: image.ImageData,
	}

	_, err := r.queries.CreatePostImage(ctx, params)
	if err != nil {
		return apperrors.WrapDatabaseError(err, "salvar imagem do post")
	}

	return nil
}

func (r *PostImageRepositoryImpl) GetPostImages(postID string) ([]*entity.PostImage, error) {
	ctx := context.Background()

	imagesDB, err := r.queries.GetPostImages(ctx, postID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar imagens do post")
	}

	images := make([]*entity.PostImage, len(imagesDB))
	for i, imgDB := range imagesDB {
		images[i] = &entity.PostImage{
			ID:        imgDB.ID,
			PostID:    imgDB.PostID,
			ImageData: imgDB.ImageData,
			CreatedAt: imgDB.CreatedAt.Time,
		}
	}

	return images, nil
}

func (r *PostImageRepositoryImpl) DeletePostImagesByPostID(postID string) error {
	ctx := context.Background()

	if err := r.queries.DeletePostImagesByPostID(ctx, postID); err != nil {
		return apperrors.WrapDatabaseError(err, "deletar imagens do post")
	}

	return nil
}

func (r *PostImageRepositoryImpl) DeletePostImageByID(imageID string) error {
	ctx := context.Background()

	if err := r.queries.DeletePostImageByID(ctx, imageID); err != nil {
		return apperrors.WrapDatabaseError(err, "deletar imagem")
	}

	return nil
}
