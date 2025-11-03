//go:build wireinject
// +build wireinject

// filepath: /Users/joaoremonato/Projects/go/maconaria_back-end/internal/infra/di/wire.go
package di

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/joaolima7/maconaria_back-end/config"
	userdata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/user"
	userdomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/user_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/handlers"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/routes"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/server"
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

// WebSet agrupa providers da camada web
var WebSet = wire.NewSet(
	handlers.NewUserHandler,
	routes.NewRouter,
	provideChiRouter,
	provideServer,
)

// provideChiRouter configura o router Chi
func provideChiRouter(router *routes.Router) *chi.Mux {
	return router.Setup()
}

// provideServer cria instância do servidor
func provideServer(router *chi.Mux, cfg *config.Config) *server.Server {
	port := "8080" // pode adicionar no .env depois
	return server.NewServer(router, port)
}

// InitializeServer injeta todas as dependências e retorna o servidor configurado
func InitializeServer(cfg *config.Config) (*server.Server, func(), error) {
	wire.Build(
		InfraSet,
		UserRepositorySet,
		UserUseCaseSet,
		WebSet,
	)
	return nil, nil, nil
}
