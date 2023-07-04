package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
 // struct - model
type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// mock data
var items = map[int]string {
	1:"keyboard",
	2:"monitor",
	3:"headphone", 
}

func main() {
	// create a new router
	r := mux.NewRouter()

 // endpoints, handler functions and HTTP methods
	r.HandleFunc("/items", listItemsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnlistItemsHandler")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

