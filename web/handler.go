package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AsgerNoer/Todo-service/data"
	"github.com/AsgerNoer/Todo-service/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//NewHandler returns a Handler struct containing a mux router and data store
func NewHandler(store *data.Store) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
		store:  *store,
	}

	//Register routes
	h.HandleFunc("/", h.todoList()).Methods("GET", "DELETE")
	h.HandleFunc("/item", h.item()).Methods("POST", "GET", "PUT", "DELETE")

	return h
}

//Handler contains a mux router and storage layer
type Handler struct {
	*mux.Router
	store data.Store
}

func (h *Handler) item() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		switch r.Method {
		//Create new Todo item
		case http.MethodPost:
			var i models.Item
			jsonObject, _ := ioutil.ReadAll(r.Body)

			err := json.Unmarshal(jsonObject, &i)
			if err != nil {
				log.Println(err)
			}

			err = h.store.CreateItem(&i)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			result, err := json.Marshal(&i)
			if err != nil {
				http.Error(w, "Unable to return json", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusAccepted)
			w.Write(result)

		//Read todo item from storage based on UUID from URL parameters
		case http.MethodGet:
			key, ok := r.URL.Query()["ID"]
			if !ok {
				http.Error(w, "Missing 'ID'", http.StatusBadRequest)
				return
			}

			ID, err := uuid.Parse(key[0])
			if err != nil {
				http.Error(w, "Malformed 'ID'", http.StatusBadRequest)
				return
			}

			item, err := h.store.ReadItem(ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result, err := json.Marshal(item)
			if err != nil {
				http.Error(w, "Unable to return json", http.StatusInternalServerError)
				return
			}

			w.Write(result)

		//Update Todo item based on UUID from URL parameters
		case http.MethodPut:
			var inputItem models.Item

			key, ok := r.URL.Query()["ID"]
			if !ok {
				http.Error(w, "Missing 'ID'", http.StatusBadRequest)
				return
			}

			ID, err := uuid.Parse(key[0])
			if err != nil {
				http.Error(w, "Malformed 'ID'", http.StatusBadRequest)
				return
			}

			jsonObject, _ := ioutil.ReadAll(r.Body)
			err = json.Unmarshal(jsonObject, &inputItem)
			if err != nil {
				log.Println(err)
			}

			inputItem.ID = ID

			err = h.store.UpdateItem(&inputItem)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			result, err := json.Marshal(inputItem)
			if err != nil {
				http.Error(w, "Unable to return json", http.StatusInternalServerError)
				return
			}

			w.Write(result)

		//Delete entry in storage based on UUID
		case http.MethodDelete:
			key, ok := r.URL.Query()["ID"]
			if !ok {
				http.Error(w, "Missing 'ID'", http.StatusBadRequest)
				return
			}

			ID, err := uuid.Parse(key[0])
			if err != nil {
				http.Error(w, "Malformed 'ID'", http.StatusBadRequest)
				return
			}

			err = h.store.DeleteItem(ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func (h *Handler) todoList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		switch r.Method {
		//Read all items on todo list from storage
		case http.MethodGet:
			item, err := h.store.GetTodoList()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result, err := json.Marshal(item)
			if err != nil {
				http.Error(w, "Unable to return json", http.StatusInternalServerError)
				return
			}

			w.Write(result)

		//Delete entire list of todo items
		case http.MethodDelete:
			err := h.store.ClearTodoList()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		}
	}
}
