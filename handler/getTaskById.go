package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/todo-host/todo-host-api/entity"
)

// GetTaskByIdHandler godoc
// @Summary Get a task
// @Description Get a task by it id
// @Tags tasks
// @Security BasicAuth
// @Accept json
// @Produce json
// @Param id path int true "Task ID" Format(int64)
// @Success 200 {object} entity.Task
// @Router /task/{id} [get]
func GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	task, err := entity.GetTaskById(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(task)
}
