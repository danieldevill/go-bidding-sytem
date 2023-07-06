package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/danieldevill/go-bidding-system/user"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bidding System!"))
	})

	user.AddMockUsers()

	// Add User HTTP handlers
	router.Get("/users", user.GetUsers)
	router.Get("/users/id/{id}", user.GetUserByID)
	router.Post("/users/id/{id}/name/{name}", user.AddUser)
	router.Delete("/users/id/{id}", user.DeleteUserByID)

	// Start an HTTP server with a given address and handler
	http.ListenAndServe(":3000", router)
}
