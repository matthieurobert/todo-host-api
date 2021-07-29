package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/todo-host/todo-host-api/entity"
)

// DeleteTaskByIdHandler godoc
// @Summary Delete a task
// @Description Delete by task id
// @Tags tasks
// @Security BasicAuth
// @Accept json
// @Produce json
// @Param id path int true "Task ID" Format(int64)
// @Success 204 {object} entity.Task
// @Router /task/{id} [delete]
func DeleteTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := entity.DeleteTaskById(int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("Task deleted")
}
