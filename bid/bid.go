package bid

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danieldevill/go-bidding-system/item"
	"github.com/danieldevill/go-bidding-system/user"
	"github.com/go-chi/chi/v5"
)

type Bid struct {
	UserID string `json:"userid"`
	ItemID string `json:"itemid"`
	Amount string `json:"amount"`
}

var bids []Bid

func AddMockBids() {
	// Bidding war 1
	bids = append(bids, Bid{UserID: "54598", ItemID: "24595", Amount: "89"})
	bids = append(bids, Bid{UserID: "54597", ItemID: "24595", Amount: "21"})
	bids = append(bids, Bid{UserID: "54594", ItemID: "24595", Amount: "78"})
	// Bidding war 2
	bids = append(bids, Bid{UserID: "54596", ItemID: "24594", Amount: "98"})
	bids = append(bids, Bid{UserID: "54595", ItemID: "24594", Amount: "67"})
	bids = append(bids, Bid{UserID: "54592", ItemID: "24594", Amount: "90"})
}

func GetBids(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bids)
}

func GetBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bidUserIDParam := chi.URLParam(r, "userid")
	bidItemIDParam := chi.URLParam(r, "itemid")

	bid := FindBid(bidUserIDParam, bidItemIDParam)

	json.NewEncoder(w).Encode(bid)
}

func FindBid(userid, itemid string) Bid {
	var b Bid
	for _, bid := range bids {
		if (bid.ItemID == itemid) && (bid.UserID == userid) {
			return bid
		}
	}
	return b
}

func BidUpdate(userid, itemid, amount string) {
	for index, bid := range bids {
		if bid.UserID == userid && bid.ItemID == itemid {
			bids = append(bids[:index], bids[index+1:]...)
			var bid Bid
			bid.UserID = userid
			bid.ItemID = itemid
			bid.Amount = amount
			bids = append(bids, bid)
			return
		}
	}
}

func AddBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bidUserIDParam := chi.URLParam(r, "userid")
	bidItemIDParam := chi.URLParam(r, "itemid")
	bidAmountParam := chi.URLParam(r, "amount")

	var bid Bid
	_ = json.NewDecoder(r.Body).Decode(bid)
	bid.UserID = bidUserIDParam
	bid.ItemID = bidItemIDParam
	bid.Amount = bidAmountParam

	if item.FindItem(bid.ItemID) != (item.Item{}) && user.FindUser(bid.UserID) != (user.User{}) {
		if FindBid(bidUserIDParam, bidItemIDParam) != (Bid{}) {
			BidUpdate(bidUserIDParam, bidItemIDParam, bidAmountParam)
		} else {
			bids = append(bids, bid)
		}

		json.NewEncoder(w).Encode(&bid)
	} else {
		if item.FindItem(bid.ItemID) == (item.Item{}) {
			bid.ItemID = "Item Not Found"
		}

		if user.FindUser(bid.UserID) == (user.User{}) {
			bid.UserID = "User Not Found"
		}
		json.NewEncoder(w).Encode(&bid)

	}
}

func UpdateBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bidUserIDParam := chi.URLParam(r, "userid")
	bidItemIDParam := chi.URLParam(r, "itemid")
	bidAmountParam := chi.URLParam(r, "amount")
	BidUpdate(bidUserIDParam, bidItemIDParam, bidAmountParam)

	json.NewEncoder(w).Encode(bids)
}

func DeleteBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bidUserIDParam := chi.URLParam(r, "userid")
	bidItemIDParam := chi.URLParam(r, "itemid")
	for index, bid := range bids {
		if bid.UserID == bidUserIDParam && bid.ItemID == bidItemIDParam {
			bids = append(bids[:index], bids[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(bids)
}

func WinnerBidByItemID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bidItemIDParam := chi.URLParam(r, "itemid")

	var winner Bid
	winner.Amount = "0"
	for _, bid := range bids {
		if bid.ItemID == bidItemIDParam {
			bidAmount, _ := strconv.Atoi(bid.Amount)
			winnerAmount, _ := strconv.Atoi(winner.Amount)
			if bidAmount > winnerAmount {
				winner = bid
			}

		}
	}
	json.NewEncoder(w).Encode(winner)
}

func BidsByItemID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bidItemIDParam := chi.URLParam(r, "itemid")

	var allBids []Bid
	for _, bid := range bids {
		if bid.ItemID == bidItemIDParam {

			allBids = append(allBids, bid)

		}
	}

	json.NewEncoder(w).Encode(allBids)
}

func ItemByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bidUserIDParam := chi.URLParam(r, "userid")

	var allItems []item.Item
	for _, bid := range bids {
		if bid.UserID == bidUserIDParam {
			item := item.FindItem(bid.ItemID)
			allItems = append(allItems, item)

		}
	}
	json.NewEncoder(w).Encode(allItems)
}
