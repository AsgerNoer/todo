package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/AsgerNoer/Todo-service/models"
)

const (
	file string = "./data.json"
)

//NewStore returns a Store struck with interfaces to retrive both items and todolist
func NewStore() (*Store, error) {
	fileinfo, err := os.Stat(file)

	//Check if file exitst
	if os.IsNotExist(err) {
		if err := newDataFile(); err != nil {
			log.Fatalf("Not able to create datafile: %q", err)
		}
		log.Println("new data file added")

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

func newDataFile() error {
	newFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer newFile.Close()

	result, _ := json.Marshal(models.TodoList{})
	err = ioutil.WriteFile(newFile.Name(), result, 0077)
	if err != nil {
		return err
	}
	return nil
}

//GetTodoList takes a and returns the data store on disk in a struct
func GetTodoList(file *os.File) models.TodoList {

	todoList := models.TodoList{}

	//Open file at unmarshal the content
	jsonObject, _ := ioutil.ReadAll(file)

	err := json.Unmarshal(jsonObject, &todoList)
	if err != nil {
		log.Printf("error unmarshalling: %q, %q", err, jsonObject)
	}
	return todoList
}
