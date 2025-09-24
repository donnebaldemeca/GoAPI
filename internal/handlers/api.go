package handlers

import (
	"github.com/donnebaldemeca/GoAPI/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes) // Use() allows us to add middleware functions to the router
	// chimiddle.StripSlashes removes slashes from the URL, e.g., /path/ becomes /path

	r.Route("/account", func(router chi.Router) { // Create a new route group for /account
		router.Use(middleware.Authorization) // Apply Authorization middleware to all routes under /account
		router.Get("/coins", GetCoinBalance) // GET /account/coins endpoint to get coin balance, using GetCoinBalance handler
	})

}
