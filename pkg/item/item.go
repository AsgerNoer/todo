package item

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Added    time.Time `json:"added,omitempty"`
	Updated  time.Time `json:"updated,omitempty"`
	ItemText string    `json:"itemText,omitempty"`
}

func (i *Item) NewItem() {
	currentTime := time.Now()
	i.ID = uuid.New()
	i.Added = currentTime
	i.Updated = currentTime
}

type Repository interface {
	ReadAllItems() ([]Item, error)
	ReadItem(id string) (Item, error)
	CreateItem(input string) (Item, error)
	UpdateItem(id, text string) error
	DeleteItem(id string) error
}

type Service interface {
	ReadAllItems() ([]Item, error)
	ReadItem(ids ...string) ([]Item, error)
	CreateItem(inputs ...string) ([]Item, error)
	UpdateItem(id, text string) error
	DeleteItem(ids ...string) error
}

type service struct {
	repo Repository
}

func NewItemService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) ReadAllItems() ([]Item, error) {
	return s.repo.ReadAllItems()
}

func (s *service) ReadItem(ids ...string) ([]Item, error) {
	var items []Item

	for _, id := range ids {
		item, err := s.repo.ReadItem(id)
		if err != nil {
			return nil, fmt.Errorf("reading id: %s from store caused an error: %v", id, err)
		}

		items = append(items, item)
	}

	return items, nil
}

func (s *service) CreateItem(inputs ...string) ([]Item, error) {
	var items []Item

	for _, input := range inputs {
		item, err := s.repo.CreateItem(input)
		if err != nil {
			return nil, fmt.Errorf("adding '%s' to store caused an error: %v", input, err)
		}

		items = append(items, item)
	}

	return items, nil
}

func (s *service) UpdateItem(id string, text string) error {
	return s.repo.UpdateItem(id, text)
}

func (s *service) DeleteItem(ids ...string) error {
	for _, id := range ids {
		if err := s.repo.DeleteItem(id); err != nil {
			return fmt.Errorf("unable to delete '%s' from store caused an error: %v", id, err)
		}
	}

	return nil
}
