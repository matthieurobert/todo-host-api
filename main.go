package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/todo-host/todo-host-api/auth"
	"github.com/todo-host/todo-host-api/config"
	"github.com/todo-host/todo-host-api/entity"
	"github.com/todo-host/todo-host-api/handler"
	"github.com/todo-host/todo-host-api/middleware"
)

func main() {
	config.InitConfig()
	auth.InitAuth()
	entity.CreateSchema(config.POSTGRES.DB)

	router := mux.NewRouter()

	router.HandleFunc("/v1/auth/token", middleware.AuthMiddleware(http.HandlerFunc(auth.CreateToken))).Methods("GET")

	router.HandleFunc("/v1/task", middleware.AuthMiddleware(http.HandlerFunc(handler.CreateTaskHandler))).Methods("POST")
	router.HandleFunc("/v1/tasks", middleware.AuthMiddleware(http.HandlerFunc(handler.GetTasksHandler))).Methods("GET")
	router.HandleFunc("/v1/task/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetTaskByIdHandler))).Methods("GET")
	router.HandleFunc("/v1/task/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.DeleteTaskByIdHandler))).Methods("DELETE")

	router.HandleFunc("/v1/register", handler.Register).Methods("POST")

	http.ListenAndServe(":8000", router)
}
