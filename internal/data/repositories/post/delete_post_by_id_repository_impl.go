package post

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type DeletePostRepositoryImpl struct {
	queries *db.Queries
}

func NewDeletePostRepositoryImpl(queries *db.Queries) *DeletePostRepositoryImpl {
	return &DeletePostRepositoryImpl{queries: queries}
}

func (r *DeletePostRepositoryImpl) Delete(postID string) error {
	ctx := context.Background()

	_, err := r.queries.GetPostByID(ctx, postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.NewNotFoundError("Post")
		}
		return apperrors.WrapDatabaseError(err, "buscar post")
	}

	if err := r.queries.DeletePost(ctx, postID); err != nil {
		return apperrors.WrapDatabaseError(err, "deletar post")
	}

	return nil
}
