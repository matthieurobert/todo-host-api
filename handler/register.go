package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/todo-host/todo-host-api/entity"
)

// Register godoc
// @Summary Register a new user
// @Description Create a new user with the input payload
// @Tags auth
// @Accept json
// @Produce json
// @Param user body entity.User true "Create User"
// @Success 200 {object} entity.User
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	// username := r.FormValue("username")
	// password := r.FormValue("password")
	// email := r.FormValue("email")

	// if username == "" || password == "" || email == "" {
	// 	http.Error(w, "missed data", http.StatusInternalServerError)
	// }

	// id, err := entity.PostUser(username, password, email)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)

	// 	json.NewEncoder(w).Encode(id)
	// }

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var user entity.User

	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if user.Email == "" || user.Username == "" || user.Password == "" {
		http.Error(w, "missed data", http.StatusInternalServerError)
	}

	id, err := entity.PostUser(user.Username, user.Password, user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(id)
}
