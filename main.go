package main

import (
	"log"
	"net/http"

	"github.com/AsgerNoer/Todo-service/data"
	"github.com/AsgerNoer/Todo-service/web"
)

func main() {
	log.Println("Service starting")

	store, err := data.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	handler := web.NewHandler(store)

	log.Fatal(http.ListenAndServe(":3000", handler.Router))
}
