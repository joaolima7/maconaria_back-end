package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/handlers"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/middlewares"
)

type Router struct {
	UserHandler     *handlers.UserHandler
	AuthHandler     *handlers.AuthHandler
	PostHandler     *handlers.PostHandler
	WorkerHandler   *handlers.WorkerHandler
	TimelineHandler *handlers.TimelineHandler
	AcaciaHandler   *handlers.AcaciaHandler
	LibraryHandler  *handlers.LibraryHandler
	HealthHandler   *handlers.HealthHandler
	AuthMiddleware  *middlewares.AuthMiddleware
}

func NewRouter(
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
	postHandler *handlers.PostHandler,
	workerHandler *handlers.WorkerHandler,
	timelineHandler *handlers.TimelineHandler,
	acaciaHandler *handlers.AcaciaHandler,
	libraryHandler *handlers.LibraryHandler,
	healthHandler *handlers.HealthHandler,
	authMiddleware *middlewares.AuthMiddleware,
) *Router {
	return &Router{
		UserHandler:     userHandler,
		AuthHandler:     authHandler,
		PostHandler:     postHandler,
		WorkerHandler:   workerHandler,
		TimelineHandler: timelineHandler,
		AcaciaHandler:   acaciaHandler,
		LibraryHandler:  libraryHandler,
		HealthHandler:   healthHandler,
		AuthMiddleware:  authMiddleware,
	}
}

func (rt *Router) Setup() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/health", rt.HealthHandler.Check)
	r.Head("/health", rt.HealthHandler.Check)

	r.Route("/api", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Route("/auth", func(r chi.Router) {
				r.Post("/login", rt.AuthHandler.Login)
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(rt.AuthMiddleware.Authenticate)

			r.Route("/users", func(r chi.Router) {
				r.Get("/", rt.UserHandler.GetAllUsers)
				r.Get("/{id}", rt.UserHandler.GetUserById)
				r.Post("/", rt.UserHandler.CreateUser)
				r.Put("/{id}", rt.UserHandler.UpdateUser)
				r.Patch("/{id}/password", rt.UserHandler.UpdateUserPassword)
			})

			r.Route("/posts", func(r chi.Router) {
				r.Post("/", rt.PostHandler.CreatePost)
				r.Get("/", rt.PostHandler.GetAllPosts)
				r.Put("/{id}", rt.PostHandler.UpdatePost)
				r.Delete("/{id}", rt.PostHandler.DeletePost)
			})

			r.Route("/workers", func(r chi.Router) {
				r.Post("/", rt.WorkerHandler.CreateWorker)
				r.Get("/", rt.WorkerHandler.GetAllWorkers)
				r.Get("/{id}", rt.WorkerHandler.GetWorkerByID)
				r.Put("/{id}", rt.WorkerHandler.UpdateWorker)
				r.Delete("/{id}", rt.WorkerHandler.DeleteWorker)
			})

			r.Route("/timelines", func(r chi.Router) {
				r.Post("/", rt.TimelineHandler.CreateTimeline)
				r.Get("/", rt.TimelineHandler.GetAllTimelines)
				r.Get("/{id}", rt.TimelineHandler.GetTimelineByID)
				r.Put("/{id}", rt.TimelineHandler.UpdateTimeline)
				r.Delete("/{id}", rt.TimelineHandler.DeleteTimeline)
			})

			r.Route("/acacias", func(r chi.Router) {
				r.Post("/", rt.AcaciaHandler.CreateAcacia)
				r.Get("/", rt.AcaciaHandler.GetAllAcacias)
				r.Get("/{id}", rt.AcaciaHandler.GetAcaciaByID)
				r.Put("/{id}", rt.AcaciaHandler.UpdateAcacia)
				r.Delete("/{id}", rt.AcaciaHandler.DeleteAcacia)
			})

			r.Route("/libraries", func(r chi.Router) {
				r.Post("/", rt.LibraryHandler.CreateLibrary)
				r.Get("/", rt.LibraryHandler.GetAllLibraries)
				r.Get("/{id}", rt.LibraryHandler.GetLibraryByID)
				r.Get("/degree/{degree}", rt.LibraryHandler.GetLibrariesByDegree)
				r.Put("/{id}", rt.LibraryHandler.UpdateLibrary)
				r.Delete("/{id}", rt.LibraryHandler.DeleteLibrary)
			})
		})
	})

	return r
}
