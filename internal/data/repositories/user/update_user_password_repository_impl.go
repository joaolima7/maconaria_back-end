package user

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type UpdateUserPasswordRepositoryImpl struct {
	queries *db.Queries
}

func NewUpdateUserPasswordRepositoryImpl(queries *db.Queries) *UpdateUserPasswordRepositoryImpl {
	return &UpdateUserPasswordRepositoryImpl{
		queries: queries,
	}
}

func (r *UpdateUserPasswordRepositoryImpl) UpdateUserPassword(userID string, newPassword string) error {
	ctx := context.Background()

	params := db.UpdateUserPasswordParams{
		ID:       userID,
		Password: newPassword,
	}

	return r.queries.UpdateUserPassword(ctx, params)
}
