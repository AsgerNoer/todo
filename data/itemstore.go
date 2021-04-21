package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/AsgerNoer/Todo-service/models"
	"github.com/google/uuid"
)

func NewItemStore(file string) *ItemStore {
	return &ItemStore{
		File: file,
	}
}

type ItemStore struct {
	File string
}

func (s *ItemStore) CreateItem(item *models.Item) error {

	file, err := os.Open(s.File)
	if err != nil {
		log.Printf("error opening file: %q", err)
	}
	defer file.Close()

	todoList := GetTodoList(file)

	//Append new file to list struct
	item.NewItem()
	todoList.Items = append(todoList.Items, *item)

	//Overwrite file again
	result, _ := json.Marshal(todoList)
	err = ioutil.WriteFile(s.File, result, 0077)
	if err != nil {
		log.Println(err)

		return err
	}

	log.Printf("Task added: %q", item.ItemText)
	return nil
}

func (s *ItemStore) ReadItem(ID uuid.UUID) (models.Item, error) {
	file, err := os.Open(s.File)
	if err != nil {
		log.Printf("error opening file: %q", err)
	}
	defer file.Close()

	todoList := GetTodoList(file)

	for i := 0; i < len(todoList.Items); i++ {
		if todoList.Items[i].ID == ID {
			log.Printf("Task found: %q", todoList.Items[i].ItemText)
			return todoList.Items[i], nil
		}
	}

	return models.Item{}, errors.New("element does not exist")
}

func (s *ItemStore) UpdateItem(item *models.Item) error {
	var itemFound bool
	var oldText string
	var newText string = item.ItemText

	file, err := os.Open(s.File)
	if err != nil {
		log.Printf("error opening file: %q", err)
		return err
	}
	defer file.Close()

	todoList := GetTodoList(file)

	for i := 0; i < len(todoList.Items); i++ {
		if todoList.Items[i].ID == item.ID {
			itemFound = true
			oldText = todoList.Items[i].ItemText
			todoList.Items[i].Updated = time.Now()
			todoList.Items[i].ItemText = item.ItemText

			*item = todoList.Items[i]

			//Overwrite storage file
			result, _ := json.Marshal(todoList)
			err = ioutil.WriteFile(s.File, result, 0077)
			if err != nil {
				log.Println(err)

				return err
			}
			log.Printf("Task updated: %q updated to %q", oldText, newText)
			break
		}
	}

	if !itemFound {
		return errors.New("element do not exist")
	}

	return nil
}

func (s *ItemStore) DeleteItem(ID uuid.UUID) error {
	var itemFound bool

	file, err := os.Open(s.File)
	if err != nil {
		log.Printf("error opening file: %q", err)
		return err
	}
	defer file.Close()

	todoList := GetTodoList(file)

	for i := 0; i < len(todoList.Items); i++ {
		if todoList.Items[i].ID == ID {
			itemFound = true

			todoList.Items = append(todoList.Items[:i], todoList.Items[i+1:]...)

			//Overwrite storage file
			result, _ := json.Marshal(todoList)
			err = ioutil.WriteFile(s.File, result, 0077)
			if err != nil {
				log.Println(err)

				return err
			}
			log.Printf("Task removed: %q", ID)
			break
		}
	}

	if !itemFound {
		return errors.New("element do not exist")
	}

	return nil
}
