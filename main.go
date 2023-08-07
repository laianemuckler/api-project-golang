package main

import (
	//"fmt"
	"log"
)

func main() {
	database, err := PostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Init(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println("%+v\n", database)

	server := NewAPIServer(":8000", database)
	server.Run()
}

// import (
// 	//"database/sql"
// 	"encoding/json"
// 	//"fmt"
// 	"log"
// 	"net/http"
// 	//"os"
// 	"strconv"

// 	"github.com/gorilla/mux"
// 	//"github.com/lib/pq"
// )
// //struct - model

// type Item struct {
// 	Id   int    `json:"id"` // metadata - id key of JSON will map to the Id field of the struct
// 	Name string `json:"name"` // capital letter so can be export and parse
// }

// var items []Item

// func initData() []Item {
// 	items = []Item {
// 		{ Id: 1, Name: "keyboard"}, { Id: 2, Name: "monitor"}, { Id: 3, Name: "headphone"},
// 	}
// 	return items
// }

// func main() {
// 	    //connect to database
// 			// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
// 			// if err != nil {
// 			// 	panic(err)
// 			// }
// 			// defer db.Close()

// 	// create a new router
// 	r := mux.NewRouter()

// 	items = append(initData())

//  // endpoints, handler functions and HTTP methods
// 	r.HandleFunc("/items", listItemsHandler).Methods("GET")
// 	r.HandleFunc("/items", createItemHandler).Methods("POST")
// 	r.HandleFunc("/items/{id}", updateItemHandler).Methods("PUT")
// 	r.HandleFunc("/items/{id}", deleteItemHandler).Methods("DELETE")
// 	r.HandleFunc("/items/{id}", listItemHandler).Methods("GET")

// 	// não utilizar fatal
// 	log.Fatal(http.ListenAndServe(":8000", r))
// }

// func listItemsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(items)
// }

// func listItemHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var param = mux.Vars(r)["id"]
// 	id, err  := strconv.Atoi(param)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	for _, item := range items {
// 		if  item.Id == id {
// 			json.NewEncoder(w).Encode(item)
// 		} else {
// 			w.WriteHeader(http.StatusNotFound)
// 		}
//   }
// }

// func createItemHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// declare a new Item struct
// 	var item Item
//   err := json.NewDecoder(r.Body).Decode(&item)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	items = append(items, item)

// 	json.NewEncoder(w).Encode(item)
// }

//  func updateItemHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var item Item
// 	err := json.NewDecoder(r.Body).Decode(&item)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	for index, structs := range items {
// 		if structs.Name == item.Name {
// 			items = append(items[:index], items[index+1:]...)
// 			} else {
// 				w.WriteHeader(http.StatusNotFound)
// 			}
// 	}
// 	items = append(items, item)
// 	json.NewEncoder(w).Encode(&item)
// }

// func deleteItemHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var param = mux.Vars(r)["id"]
// 	id, err  := strconv.Atoi(param)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	indexId := 0
// 	for index, structs := range items {
// 		if structs.Id == id {
// 			indexId = index
// 		} else {
// 			w.WriteHeader(http.StatusNotFound)
// 		}
// 	}
// 	// append - remove an element from a slice by creating a new slice with all elements except the one you want to remove
// 	items = append(items[:indexId], items[indexId+1:]...)
// }
