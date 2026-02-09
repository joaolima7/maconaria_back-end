package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetUserByEmailRepositoryImpl struct {
	queries *db.Queries
}

func NewGetUserByEmailRepositoryImpl(queries *db.Queries) *GetUserByEmailRepositoryImpl {
	return &GetUserByEmailRepositoryImpl{queries: queries}
}

func (r *GetUserByEmailRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	ctx := context.Background()

	userDB, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
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
