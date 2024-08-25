package handlers

import (
	"github.com/brianleogoldman/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware (applied to all endpoints we make)
	r.Use(chimiddle.StripSlashes) // StripSlashes function makes sure trailing slashes will always be ignored

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)

		// GET endpoint /coins will be handled by the GetCoinBalance function
		router.Get("/coins", GetCoinBalance)
	})
}
