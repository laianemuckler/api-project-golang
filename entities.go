package main

import "math/rand"

type CreateItemRequest struct {
	Name string `json:"name"`
}

type Item struct {
	Id   int    `json:"id"`   // metadata - id key of JSON will map to the Id field of the struct
	Name string `json:"name"` // capital letter so can be export and parse
}

func NewItem(id int, name string) *Item {
	return &Item{
		Id:   rand.Intn(10000),
		Name: name,
	}
}