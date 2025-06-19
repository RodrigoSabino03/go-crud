package main

import (
	"log"
	"net/http"

	"github.com/RodrigoSabino03/simple-go-mod/config"
	"github.com/RodrigoSabino03/simple-go-mod/handlers"
	"github.com/RodrigoSabino03/simple-go-mod/models"
	"github.com/gorilla/mux"
)

func main() {
	dbconnection := config.SetupDataBase()

	_, err := dbconnection.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(dbconnection)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")

	defer dbconnection.Close()

	log.Fatal(http.ListenAndServe(":8080", router))
}
