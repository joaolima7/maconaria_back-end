package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetUserByIdRepositoryImpl struct {
	queries *db.Queries
}

func NewGetUserByIdRepositoryImpl(queries *db.Queries) *GetUserByIdRepositoryImpl {
	return &GetUserByIdRepositoryImpl{
		queries: queries,
	}
}

func (r *GetUserByIdRepositoryImpl) GetUserById(id string) (*entity.User, error) {
	ctx := context.Background()

	userDB, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Usuário")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar usuário")
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
