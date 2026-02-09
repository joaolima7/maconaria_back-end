package user

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetAllUsersRepositoryImpl struct {
	queries *db.Queries
}

func NewGetAllUsersRepositoryImpl(queries *db.Queries) *GetAllUsersRepositoryImpl {
	return &GetAllUsersRepositoryImpl{queries: queries}
}

func (r *GetAllUsersRepositoryImpl) GetAllUsers() ([]*entity.User, error) {
	ctx := context.Background()

	usersDb, err := r.queries.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []*entity.User

	for _, userDB := range usersDb {
		user := &entity.User{
			ID:        userDB.ID,
			Name:      userDB.Name,
			Email:     userDB.Email,
			CIM:       userDB.Cim,
			Degree:    entity.UserDegree(userDB.Degree),
			IsActive:  userDB.IsActive,
			IsAdmin:   userDB.IsAdmin,
			IsRegular: userDB.IsRegular,
		}

		users = append(users, user)
	}

	return users, nil
}
