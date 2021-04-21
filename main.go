package main

import (
	"log"
	"net/http"

	"github.com/AsgerNoer/Todo-service/data"
	"github.com/AsgerNoer/Todo-service/web"
)

func main() {
	log.Println("Service starting...")

	//Init data store
	store, err := data.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	//Init handler with router
	handler := web.NewHandler(store)

	log.Println("...Service ready!")
	log.Fatal(http.ListenAndServe(":3000", handler.Router))
}
