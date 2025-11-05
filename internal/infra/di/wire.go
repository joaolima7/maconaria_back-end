//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/joaolima7/maconaria_back-end/config"
	postdata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/post"
	timelinedata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/timeline"
	userdata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/user"
	workerdata "github.com/joaolima7/maconaria_back-end/internal/data/repositories/worker"
	postdomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/post"
	timelinedomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/timeline"
	userdomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
	workerdomain "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/worker"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/post_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/timeline_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/user_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/worker_usecase"
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
	postdata.NewPostImageRepositoryImpl,
	wire.Bind(new(postdomain.PostImageRepository), new(*postdata.PostImageRepositoryImpl)),

	postdata.NewCreatePostRepositoryImpl,
	wire.Bind(new(postdomain.CreatePostRepository), new(*postdata.CreatePostRepositoryImpl)),

	postdata.NewGetAllPostsRepositoryImpl,
	wire.Bind(new(postdomain.GetAllPostsRepository), new(*postdata.GetAllPostsRepositoryImpl)),

	postdata.NewUpdatePostByIDRepositoryImpl,
	wire.Bind(new(postdomain.UpdatePostByIDRepository), new(*postdata.UpdatePostByIDRepositoryImpl)),

	postdata.NewDeletePostRepositoryImpl,
	wire.Bind(new(postdomain.DeletePostRepository), new(*postdata.DeletePostRepositoryImpl)),
)

// Worker Repository Set
var WorkerRepositorySet = wire.NewSet(
	workerdata.NewCreateWorkerRepositoryImpl,
	wire.Bind(new(workerdomain.CreateWorkerRepository), new(*workerdata.CreateWorkerRepositoryImpl)),

	workerdata.NewGetAllWorkersRepositoryImpl,
	wire.Bind(new(workerdomain.GetAllWorkersRepository), new(*workerdata.GetAllWorkersRepositoryImpl)),

	workerdata.NewGetWorkerByIDRepositoryImpl,
	wire.Bind(new(workerdomain.GetWorkerByIDRepository), new(*workerdata.GetWorkerByIDRepositoryImpl)),

	workerdata.NewUpdateWorkerByIDRepositoryImpl,
	wire.Bind(new(workerdomain.UpdateWorkerByIDRepository), new(*workerdata.UpdateWorkerByIDRepositoryImpl)),

	workerdata.NewDeleteWorkerRepositoryImpl,
	wire.Bind(new(workerdomain.DeleteWorkerRepository), new(*workerdata.DeleteWorkerRepositoryImpl)),
)

// Timeline Repository Set
var TimelineRepositorySet = wire.NewSet(
	timelinedata.NewCreateTimelineRepositoryImpl,
	wire.Bind(new(timelinedomain.CreateTimelineRepository), new(*timelinedata.CreateTimelineRepositoryImpl)),

	timelinedata.NewGetAllTimelinesRepositoryImpl,
	wire.Bind(new(timelinedomain.GetAllTimelinesRepository), new(*timelinedata.GetAllTimelinesRepositoryImpl)),

	timelinedata.NewGetTimelineByIDRepositoryImpl,
	wire.Bind(new(timelinedomain.GetTimelineByIDRepository), new(*timelinedata.GetTimelineByIDRepositoryImpl)),

	timelinedata.NewUpdateTimelineByIDRepositoryImpl,
	wire.Bind(new(timelinedomain.UpdateTimelineByIDRepository), new(*timelinedata.UpdateTimelineByIDRepositoryImpl)),

	timelinedata.NewDeleteTimelineRepositoryImpl,
	wire.Bind(new(timelinedomain.DeleteTimelineRepository), new(*timelinedata.DeleteTimelineRepositoryImpl)),
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

// Worker UseCase Set
var WorkerUseCaseSet = wire.NewSet(
	worker_usecase.NewCreateWorkerUseCase,
	worker_usecase.NewGetAllWorkersUseCase,
	worker_usecase.NewGetWorkerByIDUseCase,
	worker_usecase.NewUpdateWorkerByIDUseCase,
	worker_usecase.NewDeleteWorkerUseCase,
)

// Timeline UseCase Set
var TimelineUseCaseSet = wire.NewSet(
	timeline_usecase.NewCreateTimelineUseCase,
	timeline_usecase.NewGetAllTimelinesUseCase,
	timeline_usecase.NewGetTimelineByIDUseCase,
	timeline_usecase.NewUpdateTimelineByIDUseCase,
	timeline_usecase.NewDeleteTimelineUseCase,
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
	handlers.NewWorkerHandler,
	handlers.NewTimelineHandler,
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
		WorkerRepositorySet,
		TimelineRepositorySet,
		UserUseCaseSet,
		PostUseCaseSet,
		WorkerUseCaseSet,
		TimelineUseCaseSet,
		WebSet,
		wire.Struct(new(App), "Server", "DB"),
	)
	return nil, nil
}
