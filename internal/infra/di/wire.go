//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/joaolima7/maconaria_back-end/config"
	postdata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/post"
	userdata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/user"
	postdomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/post"
	userdomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/post_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/user_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/auth"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/handlers"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/middlewares"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/routes"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/server"
)

// User Repository Set
var UserRepositorySet = wire.NewSet(
	userdata.NewCreateUserRepositoryImpl,
	wire.Bind(new(userdomain.CreateUserRepository), new(*userdata.CreateUserRepositoryImpl)),

	userdata.NewGetAllUsersRepositoryImpl,
	wire.Bind(new(userdomain.GetAllUsersRepository), new(*userdata.GetAllUsersRepositoryImpl)),

	userdata.NewGetUserByEmailRepositoryImpl,
	wire.Bind(new(userdomain.GetUserByEmailRepository), new(*userdata.GetUserByEmailRepositoryImpl)),

	userdata.NewGetUserByIdRepositoryImpl,
	wire.Bind(new(userdomain.GetUserByIdRepository), new(*userdata.GetUserByIdRepositoryImpl)),

	userdata.NewUpdateUserByIDRepositoryImpl,
	wire.Bind(new(userdomain.UpdateUserByIDRepository), new(*userdata.UpdateUserByIDRepositoryImpl)),

	userdata.NewUpdateUserPasswordRepositoryImpl,
	wire.Bind(new(userdomain.UpdateUserPasswordRepository), new(*userdata.UpdateUserPasswordRepositoryImpl)),
)

// Post Repository Set
var PostRepositorySet = wire.NewSet(
	// PostImage Repository (usado por outros)
	postdata.NewPostImageRepositoryImpl,
	wire.Bind(new(postdomain.PostImageRepository), new(*postdata.PostImageRepositoryImpl)),

	// Post Repositories
	postdata.NewCreatePostRepositoryImpl,
	wire.Bind(new(postdomain.CreatePostRepository), new(*postdata.CreatePostRepositoryImpl)),

	postdata.NewGetAllPostsRepositoryImpl,
	wire.Bind(new(postdomain.GetAllPostsRepository), new(*postdata.GetAllPostsRepositoryImpl)),

	postdata.NewUpdatePostByIDRepositoryImpl,
	wire.Bind(new(postdomain.UpdatePostByIDRepository), new(*postdata.UpdatePostByIDRepositoryImpl)),

	postdata.NewDeletePostRepositoryImpl,
	wire.Bind(new(postdomain.DeletePostRepository), new(*postdata.DeletePostRepositoryImpl)),
)

// User UseCase Set
var UserUseCaseSet = wire.NewSet(
	user_usecase.NewCreateUserUseCase,
	user_usecase.NewGetAllUsersUseCase,
	user_usecase.NewGetUserByIdUseCase,
	user_usecase.NewUpdateUserByIdUseCase,
	user_usecase.NewUpdateUserPasswordUseCase,
	user_usecase.NewLoginUseCase,
)

// Post UseCase Set
var PostUseCaseSet = wire.NewSet(
	post_usecase.NewCreatePostUseCase,
	post_usecase.NewGetAllPostsUseCase,
	post_usecase.NewUpdatePostByIDUseCase,
	post_usecase.NewDeletePostUseCase,
)

// Infra Set
var InfraSet = wire.NewSet(
	database.ProvideDatabase,
	database.ProvideQueries,
	provideJWTService,
)

// Web Set
var WebSet = wire.NewSet(
	handlers.NewUserHandler,
	handlers.NewAuthHandler,
	handlers.NewPostHandler,
	handlers.NewHealthHandler,
	middlewares.NewAuthMiddleware,
	routes.NewRouter,
	provideChiRouter,
	provideServer,
)

func provideJWTService(cfg *config.Config) *auth.JWTService {
	return auth.NewJWTService(cfg.JWTSecret, cfg.GetJWTDuration())
}

func provideChiRouter(router *routes.Router) *chi.Mux {
	return router.Setup()
}

func provideServer(router *chi.Mux, cfg *config.Config) *server.Server {
	return server.NewServer(router, cfg.ServerPort)
}

type App struct {
	Server *server.Server
	DB     *sql.DB
}

func (a *App) Cleanup() {
	if a.DB != nil {
		a.DB.Close()
	}
}

func InitializeApp(cfg *config.Config) (*App, error) {
	wire.Build(
		InfraSet,
		UserRepositorySet,
		PostRepositorySet,
		UserUseCaseSet,
		PostUseCaseSet,
		WebSet,
		wire.Struct(new(App), "Server", "DB"),
	)
	return nil, nil
}
