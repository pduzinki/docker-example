package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-in-docker-example/controllers"
	"go-in-docker-example/models"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello there!")
}

func main() {
	// config
	dialect := "postgres"
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", "postgres", 5432, "postgres", "123", "go_in_docker_dev")

	// message service
	ms := models.NewMessageService(dialect, connectionInfo)

	// controller
	mc := controllers.NewMessageController(ms)

	// router
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/message", mc.WriteMessage).Methods("POST")
	r.HandleFunc("/messages", mc.ReadMessages).Methods("GET")

	// listen
	log.Println("now working on port :3000")
	http.ListenAndServe(":3000", r)
}
