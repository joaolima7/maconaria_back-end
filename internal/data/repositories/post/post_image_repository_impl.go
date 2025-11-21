package post

import (
	"context"
	"database/sql"
	"errors"

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
		ID:       image.ID,
		PostID:   image.PostID,
		ImageUrl: image.ImageURL,
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
			ImageURL:  imgDB.ImageUrl,
			CreatedAt: imgDB.CreatedAt.Time,
		}
	}

	return images, nil
}

func (r *PostImageRepositoryImpl) GetPostImageByID(imageID string) (*entity.PostImage, error) {
	ctx := context.Background()

	imgDB, err := r.queries.GetPostImageByID(ctx, imageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Imagem")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar imagem")
	}

	return &entity.PostImage{
		ID:        imgDB.ID,
		PostID:    imgDB.PostID,
		ImageURL:  imgDB.ImageUrl,
		CreatedAt: imgDB.CreatedAt.Time,
	}, nil
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
