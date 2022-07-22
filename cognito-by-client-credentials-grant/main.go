package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

// Get All User
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := []User{
		{
			ID:       1,
			Name:     "名前１",
			Age:      11,
			Password: "password1",
		},
		{
			ID:       2,
			Name:     "名前２",
			Age:      21,
			Password: "password2",
		},
	}

	json.NewEncoder(w).Encode(users)
}

func main() {
	// ルーターのイニシャライズ
	r := mux.NewRouter()

	// ルート(エンドポイント)
	r.HandleFunc("/api/users", getUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
