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
		Name:     user.Name,
		Email:    user.Email,
		Cim:      user.CIM,
		Degree:   db.UsersDegree(user.Degree),
		IsActive: user.IsActive,
		IsAdmin:  user.IsAdmin,
		ID:       user.ID,
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
		CIM:      userUpdated.Cim,
		Degree:   entity.UserDegree(userUpdated.Degree),
		IsActive: userUpdated.IsActive,
		IsAdmin:  userUpdated.IsAdmin,
	}, nil
}
