package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/AsgerNoer/Todo-service/models"
)

//NewStore returns a Store struck with interfaces to retrive both items and todolist
func NewStore() (*Store, error) {
	var file = "./data.json"

	fileinfo, err := os.Stat(file)
	if os.IsNotExist(err) {
		newFile, err := os.Create(file)
		if err != nil {
			log.Panicf("file cannot be created: %q", err)
		}
		defer newFile.Close()

		log.Println("new data file added")

		GetTodoList(newFile)
	} else {
		log.Printf("datafile found: %q", fileinfo.Name())
	}

	return &Store{
		ItemStore:     NewItemStore(file),
		TodoListStore: NewTodoListStore(file),
	}, nil
}

//Store contains interfaces for items and todolist
type Store struct {
	models.ItemStore
	models.TodoListStore
}

//GetTodoList takes a and returns the data store on disk in a struct
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
