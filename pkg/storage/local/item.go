package local

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/AsgerNoer/Todo-service/pkg/item"
	"github.com/google/uuid"
)

type ItemDBModel struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Added    time.Time `json:"added,omitempty"`
	Updated  time.Time `json:"updated,omitempty"`
	ItemText string    `json:"itemText,omitempty"`
}

//ReadItem reads an item from the todo list on disk base on the UUID.
func (s *Store) ReadAllItems() ([]item.Item, error) {
	file, err := os.Open(s.File)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	todoList, err := GetTodoList(file)
	if err != nil {
		return nil, fmt.Errorf("reading current todolist failed: %v", err)
	}

	var newList []item.Item
	for _, listItem := range todoList {
		newList = append(newList, item.Item(listItem))
	}

	return newList, nil
}

//ReadItem reads an item from the todo list on disk base on the UUID.
func (s *Store) ReadItem(id string) (item.Item, error) {
	file, err := os.Open(s.File)
	if err != nil {
		return item.Item{}, err
	}

	defer file.Close()

	todoList, err := GetTodoList(file)
	if err != nil {
		return item.Item{}, fmt.Errorf("reading current todolist failed: %v", err)
	}

	for _, listItem := range todoList {
		uuid, err := uuid.Parse(id)
		if err != nil {
			return item.Item{}, fmt.Errorf("cannot parse id: %v", err)
		}

		if listItem.ID == uuid {
			return item.Item(listItem), nil
		}
	}

	return item.Item{}, errors.New("element does not exist")
}

//CreateItem creates item in json file on disk
func (s *Store) CreateItem(input string) (item.Item, error) {
	file, err := os.Open(s.File)
	if err != nil {
		return item.Item{}, err
	}

	defer file.Close()

	todoList, err := GetTodoList(file)
	if err != nil {
		return item.Item{}, fmt.Errorf("reading current todolist failed: %v", err)
	}

	now := time.Now().UTC()
	newItem := ItemDBModel{
		ID:       uuid.New(),
		Added:    now,
		Updated:  now,
		ItemText: input,
	}

	todoList = append(todoList, newItem)

	//Overwrite file again
	result, _ := json.Marshal(todoList)

	err = ioutil.WriteFile(s.File, result, 0077)
	if err != nil {
		return item.Item{}, err
	}

	return item.Item(newItem), nil
}

//UpdateItem updates the item referenced. Item should contain an ID
func (s *Store) UpdateItem(id, text string) error {
	file, err := os.Open(s.File)
	if err != nil {
		return err
	}

	defer file.Close()

	todoList, err := GetTodoList(file)
	if err != nil {
		return fmt.Errorf("reading current todolist failed: %v", err)
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("cannot parse id: %v", err)
	}

	var newList []ItemDBModel
	for _, listItem := range todoList {
		if listItem.ID == uuid {
			listItem.ItemText = text
			listItem.Updated = time.Now()
		}

		newList = append(newList, listItem)
	}

	//Overwrite storage file
	result, err := json.Marshal(newList)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.File, result, 0077)
	if err != nil {
		return err
	}

	return nil
}

//DeleteItem deletes an item in the array of items contained in the todolist structure
func (s *Store) DeleteItem(id string) error {
	file, err := os.Open(s.File)
	if err != nil {
		return err
	}

	defer file.Close()

	todoList, err := GetTodoList(file)
	if err != nil {
		return fmt.Errorf("reading current todolist failed: %v", err)
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("cannot parse id: %v", err)
	}

	var newList []ItemDBModel
	for i, listItem := range todoList {
		if listItem.ID == uuid {
			todoList = append(todoList[:i], todoList[i+1:]...)
		}

		newList = append(newList, listItem)
	}

	//Overwrite storage file
	result, err := json.Marshal(newList)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.File, result, 0077)
	if err != nil {
		return err
	}

	return nil
}
