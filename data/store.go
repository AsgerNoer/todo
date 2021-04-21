package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/AsgerNoer/Todo-service/models"
)

func NewStore() (*Store, error) {
	var file = "./data.json"

	_, err := os.Create(file)

	if err != nil {
		log.Panicf("Cannot start due to datafile: %q", err)
	}

	return &Store{
		ItemStore:     NewItemStore(file),
		TodoListStore: NewTodoListStore(file),
	}, nil
}

type Store struct {
	models.ItemStore
	models.TodoListStore
}

func GetTodoList(file *os.File) models.TodoList {

	todoList := models.TodoList{}

	//Open file at unmarshal the content
	jsonObject, _ := ioutil.ReadAll(file)

	err := json.Unmarshal(jsonObject, &todoList)
	if err != nil {
		if string(jsonObject) == "" {
			log.Println("json structure not in place. Added empty todolist")
			result, _ := json.MarshalIndent(todoList, "", "\t")
			err = ioutil.WriteFile(file.Name(), result, 0077)
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Printf("error unmarshalling: %q, %q", err, jsonObject)
		}

	}
	return todoList
}
