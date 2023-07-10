package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
//struct - model
type Item struct {
	Id   string    `json:"id"`
	Name string `json:"name"`
}

var items []Item

func initData() []Item {
	items = []Item {
		Item {
			Id: "1", Name: "keyboard",
		},
		Item {
			Id: "2", Name: "monitor",
		},
		Item {
			Id: "3", Name: "headphone",
		},
	}
	return items
}

func main() {
	// create a new router
	r := mux.NewRouter()

	items = append(initData())

 // endpoints, handler functions and HTTP methods
	r.HandleFunc("/items", listItemsHandler).Methods("GET")
	r.HandleFunc("/items", createItemHandler).Methods("POST")
	r.HandleFunc("/items/{id}", updateItemHandler).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItemHandler).Methods("DELETE")
	// r.HandleFunc("/items/{id}", listItemHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnlistItemsHandler")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func updateItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, i := range items {
		if i.Id == params["id"] {
			items = append(items[:index], items[index+1:]...)
			var item Item 
			_ = json.NewDecoder(r.Body).Decode(&item)
			item.Id= params["id"]
			items = append(items, item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, i := range items {
		if i.Id == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}

