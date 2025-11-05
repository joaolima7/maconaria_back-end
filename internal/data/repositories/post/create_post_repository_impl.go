package post

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
	"github.com/joaolima7/maconaria_back-end/internal/types"
)

type CreatePostRepositoryImpl struct {
	queries         *db.Queries
	imageRepository *PostImageRepositoryImpl
}

func NewCreatePostRepositoryImpl(queries *db.Queries, imageRepo *PostImageRepositoryImpl) *CreatePostRepositoryImpl {
	return &CreatePostRepositoryImpl{
		queries:         queries,
		imageRepository: imageRepo,
	}
}

func (r *CreatePostRepositoryImpl) CreatePost(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

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

	params := db.CreatePostParams{
		ID:                  post.ID,
		Title:               post.Title,
		Category:            post.Category,
		SmallDescription:    post.SmallDescription,
		CompleteDescription: post.CompleteDescription,
		Date:                date,
		Time:                time,
		Location:            location,
		IsFeatured:          post.IsFeatured,
		PostType:            db.PostsPostType(post.PostType),
		UserID:              post.UserID,
	}

	_, err := r.queries.CreatePost(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "criar post")
	}

	if len(post.Images) > 0 {
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
		return nil, apperrors.WrapDatabaseError(err, "buscar post criado")
	}

	images, err := r.imageRepository.GetPostImages(post.ID)
	if err != nil {
		return nil, err
	}

	createdPost := dbPostToEntity(postDB)
	createdPost.Images = images

	return createdPost, nil
}

func dbPostToEntity(postDB db.Post) *entity.Post {
	var date, time, location *string
	if postDB.Date.Valid {
		date = &postDB.Date.String
	}
	if postDB.Time.Valid {
		time = &postDB.Time.String
	}
	if postDB.Location.Valid {
		location = &postDB.Location.String
	}

	return &entity.Post{
		ID:                  postDB.ID,
		Title:               postDB.Title,
		Category:            postDB.Category,
		SmallDescription:    postDB.SmallDescription,
		CompleteDescription: postDB.CompleteDescription,
		Date:                date,
		Time:                time,
		Location:            location,
		IsFeatured:          postDB.IsFeatured,
		PostType:            types.PostType(postDB.PostType),
		UserID:              postDB.UserID,
		CreatedAt:           postDB.CreatedAt.Time,
		UpdatedAt:           postDB.UpdatedAt.Time,
	}
}
