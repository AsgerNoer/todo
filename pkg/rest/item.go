package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/AsgerNoer/Todo-service/pkg/item"
)

func HandleItemCreate(service item.Service) http.HandlerFunc {
	type Request struct {
		ItemText string `json:"itemText,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var i Request
		jsonObject, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(jsonObject, &i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := service.CreateItem(i.ItemText)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Unable to return json", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write(result)
	}
}

func HandleItemReadAll(service item.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		item, err := service.ReadAllItems()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(item)
		if err != nil {
			http.Error(w, "Unable to return json", http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(result)
	}
}

func HandleItemRead(service item.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, ok := r.URL.Query()["ID"]
		if !ok {
			http.Error(w, "Missing 'ID'", http.StatusBadRequest)
			return
		}

		item, err := service.ReadItem(key...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(item)
		if err != nil {
			http.Error(w, "Unable to return json", http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(result)
	}
}

func HandleItemUpdate(service item.Service) http.HandlerFunc {
	type Request struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req Request

		jsonBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(jsonBytes, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = service.UpdateItem(req.ID, req.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(req)
		if err != nil {
			http.Error(w, "Unable to return json", http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(result)
	}
}

func HandleItemDelete(service item.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, ok := r.URL.Query()["ID"]
		if !ok {
			http.Error(w, "Missing 'ID'", http.StatusBadRequest)
			return
		}

		err := service.DeleteItem(key...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
