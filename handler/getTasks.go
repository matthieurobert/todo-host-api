package handler

import (
	"encoding/json"
	"net/http"

	"github.com/todo-host/todo-host-api/auth"
	"github.com/todo-host/todo-host-api/entity"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	_, userInfo, _ := auth.Strategy.AuthenticateRequest(r)
	user, _ := entity.GetUserByUsername(userInfo.GetUserName())

	tasks, err := entity.GetTasks(user.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(tasks)
}
