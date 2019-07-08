package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/local/TaskListGo/controllers"
	"github.com/local/TaskListGo/middleware"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	api.Use(middleware.JwtAuthentication)

	api.HandleFunc("/", controllers.Hello).Methods(http.MethodGet)
	api.HandleFunc("/todos", controllers.ListToDos).Methods(http.MethodGet)
	api.HandleFunc("/todos/{toDold}", controllers.PertoDocs).Methods(http.MethodGet)
	api.HandleFunc("/todos", controllers.CreateToDocs).Methods(http.MethodPost)
	api.HandleFunc("/todos/{toDold}/action", controllers.ActionToDocs).Methods(http.MethodPost)
	api.HandleFunc("/todos/{toDold}/edit", controllers.EditToDocs).Methods(http.MethodPost)
	api.HandleFunc("/todos/{toDold}/delete", controllers.DeleteToDocs).Methods(http.MethodPost)
	api.HandleFunc("/user/register", controllers.Register).Methods(http.MethodPost)
	api.HandleFunc("/user/login", controllers.Login).Methods(http.MethodPost)

	if e := godotenv.Load(); e != nil {
		fmt.Print(e)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	srv := &http.Server{
		Handler:      api,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
