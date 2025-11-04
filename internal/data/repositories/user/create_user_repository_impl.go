package user

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type CreateUserRepositoryImpl struct {
	queries *db.Queries
}

func NewCreateUserRepositoryImpl(queries *db.Queries) *CreateUserRepositoryImpl {
	return &CreateUserRepositoryImpl{queries: queries}
}

func (r *CreateUserRepositoryImpl) CreateUser(user *entity.User) (*entity.User, error) {
	ctx := context.Background()

	params := db.CreateUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsActive: user.IsActive,
		IsAdmin:  user.IsAdmin,
	}

	userExisting, _ := r.queries.GetUserByEmail(ctx, user.Email)
	if userExisting.Email == user.Email {
		return nil, apperrors.NewDuplicateError("e-mail", user.Email)
	}

	_, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "criar usuário")
	}

	userDB, err := r.queries.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar usuário criado")
	}

	return &entity.User{
		ID:        userDB.ID,
		Name:      userDB.Name,
		Email:     userDB.Email,
		Password:  userDB.Password,
		IsActive:  userDB.IsActive,
		IsAdmin:   userDB.IsAdmin,
		CreatedAt: userDB.CreatedAt.Time,
		UpdatedAt: userDB.UpdatedAt.Time,
	}, nil
}
