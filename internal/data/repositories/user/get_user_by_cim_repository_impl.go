package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetUserByCIMRepositoryImpl struct {
	queries *db.Queries
}

func NewGetUserByCIMRepositoryImpl(queries *db.Queries) *GetUserByCIMRepositoryImpl {
	return &GetUserByCIMRepositoryImpl{queries: queries}
}

func (r *GetUserByCIMRepositoryImpl) GetUserByCIM(cim string) (*entity.User, error) {
	ctx := context.Background()

	userDB, err := r.queries.GetUserByCIM(ctx, cim)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Usuário")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar usuário por CIM")
	}

	return &entity.User{
		ID:        userDB.ID,
		Name:      userDB.Name,
		Email:     userDB.Email,
		Password:  userDB.Password,
		CIM:       userDB.Cim,
		Degree:    entity.UserDegree(userDB.Degree),
		IsActive:  userDB.IsActive,
		IsAdmin:   userDB.IsAdmin,
		IsRegular: userDB.IsRegular,
		CreatedAt: userDB.CreatedAt.Time,
		UpdatedAt: userDB.UpdatedAt.Time,
	}, nil
}
