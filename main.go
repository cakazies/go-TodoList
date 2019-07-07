package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/local/TaskListGo/controllers"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", controllers.Hello).Methods(http.MethodGet)
	api.HandleFunc("/todos", controllers.ListToDos).Methods(http.MethodGet)
	api.HandleFunc("/todos/{toDold}", controllers.PertoDocs).Methods(http.MethodGet)
	api.HandleFunc("/todos", controllers.CreateToDocs).Methods(http.MethodPost)
	api.HandleFunc("/todos/{toDold}/action", controllers.ActionToDocs).Methods(http.MethodPost)
	api.HandleFunc("/todos/{toDold}/edit", controllers.EditToDocs).Methods(http.MethodPost)
	api.HandleFunc("/todos/{toDold}/delete", controllers.DeleteToDocs).Methods(http.MethodPost)
	api.HandleFunc("/user/register", controllers.UserRegister).Methods(http.MethodPost)
	api.HandleFunc("/user/login", controllers.UserLogin).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      api,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
