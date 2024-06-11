package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/quangd42/meal-planner/backend/internal/middleware"
)

func AddRoutes(r *chi.Mux) {
	r.Get("/", r.NotFoundHandler())

	// API router
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthz", readinessHandler)
		r.Get("/err", errorHandler)

		r.Mount("/users", usersAPIRouter())
		r.Mount("/auth", authRouter())
	})
}

// usersAPIRouter
func usersAPIRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/", createUserHandler)
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthVerifier())
		r.Put("/", updateUserHandler)
	})

	return r
}

// authRouter
func authRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", loginHandler)
	r.Post("/refresh", refreshJWTHandler)
	r.Post("/revoke", revokeJWTHandler)

	return r
}
