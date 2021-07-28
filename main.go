package main

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/todo-host/todo-host-api/docs"

	"github.com/todo-host/todo-host-api/auth"
	"github.com/todo-host/todo-host-api/config"
	"github.com/todo-host/todo-host-api/entity"
	"github.com/todo-host/todo-host-api/handler"
	"github.com/todo-host/todo-host-api/middleware"
)

// @title Todo-host API
// @version 1.0
// @description This is a sample self hosted todo list api.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /v1

// @securityDefinitions.basic BasicAuth

func main() {
	config.InitConfig()
	auth.InitAuth()
	entity.CreateSchema(config.POSTGRES.DB)

	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	router.HandleFunc("/v1/auth/token", middleware.AuthMiddleware(http.HandlerFunc(auth.CreateToken))).Methods("GET")

	router.HandleFunc("/v1/task", middleware.AuthMiddleware(http.HandlerFunc(handler.CreateTaskHandler))).Methods("POST")
	router.HandleFunc("/v1/tasks", middleware.AuthMiddleware(http.HandlerFunc(handler.GetTasksHandler))).Methods("GET")
	router.HandleFunc("/v1/task/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.GetTaskByIdHandler))).Methods("GET")
	router.HandleFunc("/v1/task/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.DeleteTaskByIdHandler))).Methods("DELETE")

	router.HandleFunc("/v1/register", handler.Register).Methods("POST")

	http.ListenAndServe(":8000", router)
}
