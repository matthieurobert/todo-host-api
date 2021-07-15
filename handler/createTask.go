package handler

import (
	"fmt"
	"net/http"

	"github.com/todo-host/todo-host-api/auth"
	"github.com/todo-host/todo-host-api/entity"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")

	_, userInfo, _ := auth.Strategy.AuthenticateRequest(r)
	user, _ := entity.GetUserByUsername(userInfo.GetUserName())

	task := entity.Task{
		Title:  title,
		Body:   body,
		UserId: user.Id,
	}

	_, err := entity.PostTask(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("Task created")
}
