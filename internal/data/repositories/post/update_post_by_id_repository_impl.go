package post

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type UpdatePostByIDRepositoryImpl struct {
	queries         *db.Queries
	imageRepository *PostImageRepositoryImpl
}

func NewUpdatePostByIDRepositoryImpl(queries *db.Queries, imageRepo *PostImageRepositoryImpl) *UpdatePostByIDRepositoryImpl {
	return &UpdatePostByIDRepositoryImpl{
		queries:         queries,
		imageRepository: imageRepo,
	}
}

func (r *UpdatePostByIDRepositoryImpl) UpdatePostByID(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	_, err := r.queries.GetPostByID(ctx, post.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Post")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar post")
	}

	var date, time, location sql.NullString
	if post.Date != nil {
		date = sql.NullString{String: *post.Date, Valid: true}
	}
	if post.Time != nil {
		time = sql.NullString{String: *post.Time, Valid: true}
	}
	if post.Location != nil {
		location = sql.NullString{String: *post.Location, Valid: true}
	}

	params := db.UpdatePostParams{
		Title:               post.Title,
		Category:            post.Category,
		SmallDescription:    post.SmallDescription,
		CompleteDescription: post.CompleteDescription,
		Date:                date,
		Time:                time,
		Location:            location,
		IsFeatured:          post.IsFeatured,
		PostType:            db.PostsPostType(post.PostType),
		ID:                  post.ID,
	}

	_, err = r.queries.UpdatePost(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "atualizar post")
	}

	if len(post.Images) > 0 {

		if err := r.imageRepository.DeletePostImagesByPostID(post.ID); err != nil {
			return nil, err
		}

		for _, image := range post.Images {
			image.PostID = post.ID
			image.ID = uuid.New().String()
			if err := r.imageRepository.CreatePostImage(image); err != nil {
				return nil, err
			}
		}
	}

	postDB, err := r.queries.GetPostByID(ctx, post.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar post atualizado")
	}

	images, err := r.imageRepository.GetPostImages(post.ID)
	if err != nil {
		return nil, err
	}

	updatedPost := dbPostToEntity(postDB)
	updatedPost.Images = images

	return updatedPost, nil
}
