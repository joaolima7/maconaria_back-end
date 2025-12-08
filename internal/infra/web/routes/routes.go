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

		r.Post("/auth/login", rt.AuthHandler.Login)

		r.Get("/posts", rt.PostHandler.GetAllPosts)
		r.Group(func(r chi.Router) {
			r.Use(rt.AuthMiddleware.Authenticate)
			r.Post("/posts", rt.PostHandler.CreatePost)
			r.Put("/posts/{id}", rt.PostHandler.UpdatePost)
			r.Delete("/posts/{id}", rt.PostHandler.DeletePost)
		})

		r.Get("/workers", rt.WorkerHandler.GetAllWorkers)
		r.Group(func(r chi.Router) {
			r.Use(rt.AuthMiddleware.Authenticate)
			r.Get("/workers/{id}", rt.WorkerHandler.GetWorkerByID)
			r.Post("/workers", rt.WorkerHandler.CreateWorker)
			r.Put("/workers/{id}", rt.WorkerHandler.UpdateWorker)
			r.Delete("/workers/{id}", rt.WorkerHandler.DeleteWorker)
		})

		r.Get("/timelines", rt.TimelineHandler.GetAllTimelines)
		r.Group(func(r chi.Router) {
			r.Use(rt.AuthMiddleware.Authenticate)
			r.Get("/timelines/{id}", rt.TimelineHandler.GetTimelineByID)
			r.Post("/timelines", rt.TimelineHandler.CreateTimeline)
			r.Put("/timelines/{id}", rt.TimelineHandler.UpdateTimeline)
			r.Delete("/timelines/{id}", rt.TimelineHandler.DeleteTimeline)
		})

		// ========== ACACIAS ==========
		r.Get("/acacias", rt.AcaciaHandler.GetAllAcacias)
		r.Group(func(r chi.Router) {
			r.Use(rt.AuthMiddleware.Authenticate)
			r.Get("/acacias/{id}", rt.AcaciaHandler.GetAcaciaByID)
			r.Post("/acacias", rt.AcaciaHandler.CreateAcacia)
			r.Put("/acacias/{id}", rt.AcaciaHandler.UpdateAcacia)
			r.Delete("/acacias/{id}", rt.AcaciaHandler.DeleteAcacia)
		})

		r.Get("/libraries", rt.LibraryHandler.GetAllLibraries)
		r.Group(func(r chi.Router) {
			r.Use(rt.AuthMiddleware.Authenticate)
			r.Get("/libraries/{id}", rt.LibraryHandler.GetLibraryByID)
			r.Get("/libraries/degree/{degree}", rt.LibraryHandler.GetLibrariesByDegree)
			r.Post("/libraries", rt.LibraryHandler.CreateLibrary)
			r.Put("/libraries/{id}", rt.LibraryHandler.UpdateLibrary)
			r.Delete("/libraries/{id}", rt.LibraryHandler.DeleteLibrary)
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
		})
	})

	return r
}
