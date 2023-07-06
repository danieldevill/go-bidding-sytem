package item

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

func AddMockItems() {
	items = append(items, Item{ID: "24598", Name: "Lot A"})
	items = append(items, Item{ID: "24597", Name: "Lot B"})
	items = append(items, Item{ID: "24596", Name: "Lot C"})
	items = append(items, Item{ID: "24595", Name: "Lot D"})
	items = append(items, Item{ID: "24594", Name: "Lot E"})
	items = append(items, Item{ID: "24593", Name: "Lot F"})
	items = append(items, Item{ID: "24592", Name: "Lot G"})
	items = append(items, Item{ID: "24591", Name: "Lot H"})
	items = append(items, Item{ID: "24590", Name: "Lot I"})
	items = append(items, Item{ID: "24589", Name: "Lot J"})
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemParam := chi.URLParam(r, "id")

	item := func() Item {
		var i Item
		for _, item := range items {
			if item.ID == itemParam {
				return item
			}
		}
		return i
	}

	json.NewEncoder(w).Encode(item())
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemID := chi.URLParam(r, "id")
	itemName := chi.URLParam(r, "name")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(item)
	item.ID = itemID
	item.Name = itemName
	items = append(items, item)
	json.NewEncoder(w).Encode(&item)
}

func DeleteItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemID := chi.URLParam(r, "id")
	for index, item := range items {
		if item.ID == itemID {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}
