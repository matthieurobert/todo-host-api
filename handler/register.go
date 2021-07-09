package handler

import (
	"encoding/json"
	"net/http"

	"github.com/todo-host/todo-host-api/entity"
)

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	if username == "" || password == "" || email == "" {
		http.Error(w, "missed data", http.StatusInternalServerError)
	}

	id, err := entity.PostUser(username, password, email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		json.NewEncoder(w).Encode(id)
	}
}
