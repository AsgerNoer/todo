package local

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const (
	file string = "./data.json"
)

//Store contains interfaces for items and todolist
type Store struct {
	File string
}

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
		File: file,
	}, nil
}

func newDataFile() error {
	todoList := []ItemDBModel{}
	testTasks := []string{"Build great stuff", "Drink beer"}

	newFile, err := os.Create(file)
	if err != nil {
		return err
	}

	defer newFile.Close()

	for _, input := range testTasks {
		now := time.Now().UTC()
		newItem := ItemDBModel{
			ID:       uuid.New(),
			Added:    now,
			Updated:  now,
			ItemText: input,
		}

		todoList = append(todoList, newItem)
	}

	result, err := json.Marshal(todoList)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(newFile.Name(), result, 0077)
	if err != nil {
		return err
	}

	return nil
}

//GetTodoList takes a and returns the data store on disk in a struct
func GetTodoList(file *os.File) ([]ItemDBModel, error) {
	var todoList []ItemDBModel

	//Open file at unmarshal the content
	jsonObject, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading: %v", err)
	}

	err = json.Unmarshal(jsonObject, &todoList)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling: %v", err)
	}

	return todoList, nil
}
