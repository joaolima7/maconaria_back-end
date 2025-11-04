package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/handlers"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/middlewares"
)

type Router struct {
	UserHandler    *handlers.UserHandler
	AuthHandler    *handlers.AuthHandler
	HealthHandler  *handlers.HealthHandler
	AuthMiddleware *middlewares.AuthMiddleware
}

func NewRouter(
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
	healthHandler *handlers.HealthHandler,
	authMiddleware *middlewares.AuthMiddleware,
) *Router {
	return &Router{
		UserHandler:    userHandler,
		AuthHandler:    authHandler,
		HealthHandler:  healthHandler,
		AuthMiddleware: authMiddleware,
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
		})
	})

	return r
}
