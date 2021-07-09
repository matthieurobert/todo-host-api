package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/todo-host/todo-host-api/config"
)

func main() {
	config.InitConfig()

	router := mux.NewRouter()

	http.ListenAndServe(":8000", router)
}
