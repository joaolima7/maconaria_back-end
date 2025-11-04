//go:build wireinject
// +build wireinject

// filepath: /Users/joaoremonato/Projects/go/maconaria_back-end/internal/infra/di/wire.go
package di

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/joaolima7/maconaria_back-end/config"
	userdata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/user"
	userdomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/user_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/auth"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/handlers"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/middlewares"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/routes"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/server"
)

// UserRepositorySet agrupa todos os providers de repositórios de user
var UserRepositorySet = wire.NewSet(
	userdata.NewCreateUserRepositoryImpl,
	wire.Bind(new(userdomain.CreateUserRepository), new(*userdata.CreateUserRepositoryImpl)),

	userdata.NewGetAllUsersRepositoryImpl,
	wire.Bind(new(userdomain.GetAllUsersRepository), new(*userdata.GetAllUsersRepositoryImpl)),

	userdata.NewGetUserByEmailRepositoryImpl,
	wire.Bind(new(userdomain.GetUserByEmailRepository), new(*userdata.GetUserByEmailRepositoryImpl)),

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
	user_usecase.NewLoginUseCase,
)

// InfraSet agrupa providers de infraestrutura
var InfraSet = wire.NewSet(
	database.ProvideDatabase,
	database.ProvideQueries,
	provideJWTService,
)

// WebSet agrupa providers da camada web
var WebSet = wire.NewSet(
	handlers.NewUserHandler,
	handlers.NewAuthHandler,
	middlewares.NewAuthMiddleware,
	routes.NewRouter,
	provideChiRouter,
	provideServer,
)

// provideJWTService cria instância do serviço JWT
func provideJWTService(cfg *config.Config) *auth.JWTService {
	return auth.NewJWTService(cfg.JWTSecret, cfg.GetJWTDuration())
}

// provideChiRouter configura o router Chi
func provideChiRouter(router *routes.Router) *chi.Mux {
	return router.Setup()
}

// provideServer cria instância do servidor
func provideServer(router *chi.Mux, cfg *config.Config) *server.Server {
	return server.NewServer(router, cfg.ServerPort)
}

// App agrupa servidor e banco para gerenciar ciclo de vida
type App struct {
	Server *server.Server
	DB     *sql.DB
}

// Cleanup fecha a conexão do banco
func (a *App) Cleanup() {
	if a.DB != nil {
		a.DB.Close()
	}
}

// InitializeApp injeta todas as dependências e retorna App com cleanup
func InitializeApp(cfg *config.Config) (*App, error) {
	wire.Build(
		InfraSet,
		UserRepositorySet,
		UserUseCaseSet,
		WebSet,
		wire.Struct(new(App), "Server", "DB"),
	)
	return nil, nil
}
