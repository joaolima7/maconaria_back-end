package user

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type UpdateUserByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewUpdateUserByIDRepositoryImpl(queries *db.Queries) *UpdateUserByIDRepositoryImpl {
	return &UpdateUserByIDRepositoryImpl{queries: queries}
}

func (r *UpdateUserByIDRepositoryImpl) UpdateUserByID(user *entity.User) (*entity.User, error) {
	ctx := context.Background()

	params := db.UpdateUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
		IsAdmin:  user.IsAdmin,
	}

	_, err := r.queries.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	userUpdated, err := r.queries.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       userUpdated.ID,
		Name:     userUpdated.Name,
		Email:    userUpdated.Email,
		IsActive: userUpdated.IsActive,
		IsAdmin:  userUpdated.IsAdmin,
	}, nil
}
