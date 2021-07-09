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
	router.HandleFunc("/v1/register", handler.Register).Methods("POST")

	http.ListenAndServe(":8000", router)
}
