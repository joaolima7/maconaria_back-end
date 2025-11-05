package post

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetAllPostsRepositoryImpl struct {
	queries         *db.Queries
	imageRepository *PostImageRepositoryImpl
}

func NewGetAllPostsRepositoryImpl(queries *db.Queries, imageRepo *PostImageRepositoryImpl) *GetAllPostsRepositoryImpl {
	return &GetAllPostsRepositoryImpl{
		queries:         queries,
		imageRepository: imageRepo,
	}
}

func (r *GetAllPostsRepositoryImpl) GetAllPosts() ([]*entity.Post, error) {
	ctx := context.Background()

	postsDB, err := r.queries.GetAllPosts(ctx)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar posts")
	}

	posts := make([]*entity.Post, len(postsDB))
	for i, postDB := range postsDB {
		post := dbPostToEntity(postDB)

		images, err := r.imageRepository.GetPostImages(post.ID)
		if err != nil {
			return nil, err
		}
		post.Images = images

		posts[i] = post
	}

	return posts, nil
}
