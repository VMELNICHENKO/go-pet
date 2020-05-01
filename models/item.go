package models

import (
	"errors"
	"fmt"
)

// Item - example of storable entity
type Item struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   uint   `json:"id"`
}

var items map[uint]*Item

//Validate item before saving
func (item *Item) Validate() (bool, error) {
	if item.Name == "" {
		return false, errors.New("Item id is required")
	}
	//TODO: add checks
	return true, nil
}

//Store Item to list
func (item *Item) Store() (*Item, error) {
	fmt.Println("Store")
	if ok, err := item.Validate(); !ok {
		return nil, err
	}

	//items[item.ID] = item

	// TODO: return only res and error. No any controller messages

	return items[item.ID], nil
}
