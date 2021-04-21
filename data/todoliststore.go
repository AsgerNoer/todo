package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/AsgerNoer/Todo-service/models"
)

func NewTodoListStore(file string) *TodoListStore {
	return &TodoListStore{
		File: file,
	}
}

type TodoListStore struct {
	File string
}

func (s *TodoListStore) GetTodoList() (models.TodoList, error) {
	file, err := os.Open(s.File)
	if err != nil {
		log.Printf("error opening file: %q", err)
	}
	defer file.Close()

	todoList := GetTodoList(file)

	return todoList, nil
}

func (s *TodoListStore) ClearTodoList() error {
	todoList := models.TodoList{}

	file, err := os.Open(s.File)
	if err != nil {
		log.Printf("error opening file: %q", err)
	}
	defer file.Close()

	//Overwrite file again
	result, _ := json.Marshal(todoList)
	err = ioutil.WriteFile(s.File, result, 0077)
	if err != nil {
		log.Println(err)

		return err
	}
	return nil
}
