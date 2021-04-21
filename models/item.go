package models

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID       uuid.UUID
	Added    time.Time
	Updated  time.Time
	ItemText string
}

type ItemStore interface {
	CreateItem(item *Item) error
	ReadItem(ID uuid.UUID) (Item, error)
	UpdateItem(item *Item) error
	DeleteItem(ID uuid.UUID) error
}

func (i *Item) NewItem() {
	currentTime := time.Now()
	i.ID = uuid.New()
	i.Added = currentTime
	i.Updated = currentTime
}
