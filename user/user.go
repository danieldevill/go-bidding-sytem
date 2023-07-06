package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users []User

func AddMockUsers() {
	users = append(users, User{ID: "54598", Name: "Brynn Keller"})
	users = append(users, User{ID: "54597", Name: "Mya House"})
	users = append(users, User{ID: "54596", Name: "Nathalie Bravo"})
	users = append(users, User{ID: "54595", Name: "Jamir McDowell"})
	users = append(users, User{ID: "54594", Name: "Paulina Dunn"})
	users = append(users, User{ID: "54593", Name: "Shepherd Fletcher"})
	users = append(users, User{ID: "54592", Name: "Dawson Burch"})
	users = append(users, User{ID: "54591", Name: "Audrey Logan"})
	users = append(users, User{ID: "54590", Name: "Andrew Blair"})
	users = append(users, User{ID: "54589", Name: "Miriam Wiggins"})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userParam := chi.URLParam(r, "id")

	user := FindUser(userParam)

	json.NewEncoder(w).Encode(user)
}

func FindUser(userParam string) User {
	var i User
	for _, user := range users {
		if user.ID == userParam {
			return user
		}
	}
	return i
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := chi.URLParam(r, "id")
	userName := chi.URLParam(r, "name")

	var user User
	_ = json.NewDecoder(r.Body).Decode(user)
	user.ID = userID
	user.Name = userName
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := chi.URLParam(r, "id")
	for index, user := range users {
		if user.ID == userID {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
