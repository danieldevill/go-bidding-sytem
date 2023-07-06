package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/danieldevill/go-bidding-system/bid"
	"github.com/danieldevill/go-bidding-system/item"
	"github.com/danieldevill/go-bidding-system/user"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bidding System!"))
	})

	user.AddMockUsers()
	item.AddMockItems()
	bid.AddMockBids()

	// Add User HTTP handlers
	router.Get("/users", user.GetUsers)
	router.Get("/users/id/{id}", user.GetUserByID)
	router.Post("/users/id/{id}/name/{name}", user.AddUser)
	router.Delete("/users/id/{id}", user.DeleteUserByID)

	// Add Item HTTP handlers
	router.Get("/items", item.GetItems)
	router.Get("/items/id/{id}", item.GetItemByID)
	router.Post("/items/id/{id}/name/{name}", item.AddItem)
	router.Delete("/items/id/{id}", item.DeleteItemByID)

	// Add Bid HTTP handlers
	router.Get("/bids", bid.GetBids)
	router.Get("/bids/{userid}/{itemid}", bid.GetBid)
	router.Post("/bids/{userid}/{itemid}/{amount}", bid.AddBid)
	router.Put("/bids/{userid}/{itemid}/{amount}", bid.UpdateBid)
	router.Delete("/bids/{userid}/{itemid}", bid.DeleteBid)
	router.Get("/winner/{itemid}", bid.WinnerBidByItemID)
	router.Get("/bids/{itemid}", bid.BidsByItemID)
	router.Get("/items/user/{userid}", bid.ItemByUserID)

	// Start an HTTP server with a given address and handler
	http.ListenAndServe(":3000", router)
}
