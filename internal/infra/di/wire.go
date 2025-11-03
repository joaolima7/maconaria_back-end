//go:build wireinject
// +build wireinject

// filepath: /Users/joaoremonato/Projects/go/maconaria_back-end/internal/infra/di/wire.go
package di

import (
	"github.com/google/wire"
	"github.com/joaolima7/maconaria_back-end/config"
	userdata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/user"
	userdomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/user_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database"
)

// UserRepositorySet agrupa todos os providers de repositórios de user
var UserRepositorySet = wire.NewSet(
	userdata.NewCreateUserRepositoryImpl,
	wire.Bind(new(userdomain.CreateUserRepository), new(*userdata.CreateUserRepositoryImpl)),

	userdata.NewGetAllUsersRepositoryImpl,
	wire.Bind(new(userdomain.GetAllUsersRepository), new(*userdata.GetAllUsersRepositoryImpl)),

	userdata.NewUpdateUserByIDRepositoryImpl,
	wire.Bind(new(userdomain.UpdateUserByIDRepository), new(*userdata.UpdateUserByIDRepositoryImpl)),

	userdata.NewUpdateUserPasswordRepositoryImpl,
	wire.Bind(new(userdomain.UpdateUserPasswordRepository), new(*userdata.UpdateUserPasswordRepositoryImpl)),
)

// UserUseCaseSet agrupa todos os use cases de user
var UserUseCaseSet = wire.NewSet(
	user_usecase.NewCreateUserUseCase,
	user_usecase.NewGetAllUsersUseCase,
	user_usecase.NewUpdateUserByIdUseCase,
	user_usecase.NewUpdateUserPasswordUseCase,
)

// InfraSet agrupa providers de infraestrutura
var InfraSet = wire.NewSet(
	database.ProvideDatabase,
	database.ProvideQueries,
)

// UserUseCases agrupa todos os use cases em uma struct
type UserUseCases struct {
	CreateUser         *user_usecase.CreateUserUseCase
	GetAllUsers        *user_usecase.GetAllUsersUseCase
	UpdateUserById     *user_usecase.UpdateUserByIdUseCase
	UpdateUserPassword *user_usecase.UpdateUserPasswordUseCase
}

// InitializeUserUseCases injeta todas as dependências dos use cases de user
func InitializeUserUseCases(cfg *config.Config) (*UserUseCases, func(), error) {
	wire.Build(
		InfraSet,
		UserRepositorySet,
		UserUseCaseSet,
		wire.Struct(new(UserUseCases), "*"),
	)
	return nil, nil, nil
}
