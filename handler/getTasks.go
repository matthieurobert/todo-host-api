package handler

import (
	"encoding/json"
	"net/http"

	"github.com/todo-host/todo-host-api/auth"
	"github.com/todo-host/todo-host-api/entity"
)

// GetTasksHandler godoc
// @Summary Get task list
// @Description Get list of all current user tasks
// @Tags tasks
// @Security BasicAuth
// @Accept json
// @Produce json
// @Success 200 {object} entity.Task
// @Router /tasks [get]
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	_, userInfo, _ := auth.Strategy.AuthenticateRequest(r)
	user, _ := entity.GetUserByUsername(userInfo.GetUserName())

	tasks, err := entity.GetTasks(user.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(tasks)
}
