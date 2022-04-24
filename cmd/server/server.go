package main

import (
	"log"
	"net/http"

	"github.com/AsgerNoer/Todo-service/pkg/rest"
	"github.com/AsgerNoer/Todo-service/pkg/storage/local"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Service starting...")

	//Init data store
	store, err := local.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	rest.RegisterRoutes(router, store)

	log.Println("...Service ready!")
	log.Fatal(http.ListenAndServe(":3000", router))
}
