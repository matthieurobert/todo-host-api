package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/todo-host/todo-host-api/auth"
	"github.com/todo-host/todo-host-api/entity"
)

// CreateTaskHandler godoc
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Security BasicAuth
// @Accept json
// @Produce json
// @Param task body entity.Task true "Create Task"
// @Success 200 {object} entity.Task
// @Router /task [post]
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	_, userInfo, _ := auth.Strategy.AuthenticateRequest(r)
	user, _ := entity.GetUserByUsername(userInfo.GetUserName())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var task entity.Task

	err = json.Unmarshal(body, &task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	task.UserId = user.Id

	if task.Title == "" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = entity.PostTask(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("Task created")
}
