package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/handlers"
)

type Router struct {
	UserHandler *handlers.UserHandler
}

func NewRouter(userHandler *handlers.UserHandler) *Router {
	return &Router{UserHandler: userHandler}
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

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", rt.UserHandler.CreateUser)
			r.Get("/", rt.UserHandler.GetAllUsers)
			r.Put("/{id}", rt.UserHandler.UpdateUser)
			r.Patch("/{id}/password", rt.UserHandler.UpdateUserPassword)
		})
	})

	return r
}
