package rest

import (
	"github.com/AsgerNoer/Todo-service/pkg/item"
	"github.com/AsgerNoer/Todo-service/pkg/storage/local"
	"github.com/gorilla/mux"
)

//RegisterRoutes registers all routes on the router
func RegisterRoutes(r *mux.Router, store *local.Store) {
	item := item.NewItemService(store)

	r.Path("/").Methods("GET").Handler(HandleItemReadAll(item))

	r.Path("/item").Methods("GET").Handler(HandleItemRead(item))
	r.Path("/item").Methods("POST").Handler(HandleItemCreate(item))
	r.Path("/item").Methods("PUT").Handler(HandleItemUpdate(item))
	r.Path("/item").Methods("DELETE").Handler(HandleItemDelete(item))

}
