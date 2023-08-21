package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type APIServer struct {
	listenAddress string
	database Database
}

func NewAPIServer(listenAddress string, database Database) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		database: database,
	}
}

func (s *APIServer) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/items", makeHTTPHandleFunc(s.listItemsHandler)).Methods("GET")
	r.HandleFunc("/item/{id}", makeHTTPHandleFunc(s.listItemHandler)).Methods("GET")
	r.HandleFunc("/items", makeHTTPHandleFunc(s.createItemHandler)).Methods("POST")
	r.HandleFunc("/item/{id}", makeHTTPHandleFunc(s.updateItemHandler)).Methods("PUT")
	r.HandleFunc("/item/{id}", makeHTTPHandleFunc(s.deleteItemHandler)).Methods("DELETE")

	http.ListenAndServe(s.listenAddress, r)
}

func (s *APIServer) listItemsHandler(w http.ResponseWriter, r *http.Request) error {
	items, err := s.database.ListItems()
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}

	return WriteJSON(w, http.StatusOK, items)
}

func (s *APIServer) createItemHandler(w http.ResponseWriter, r *http.Request) error {
	createItemRequest := new(CreateItemRequest)
	if err := json.NewDecoder(r.Body).Decode(createItemRequest); err!= nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}

	item := NewItem(createItemRequest.Name)
	if err := s.database.CreateItem(item); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}

	return WriteJSON(w, http.StatusOK, item)
}

func (s *APIServer) updateItemHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, err.Error())
	}

	updateItemRequest := new(CreateItemRequest)
	if err := json.NewDecoder(r.Body).Decode(updateItemRequest); err!= nil {
		return WriteJSON(w, http.StatusBadRequest, "Invalid request payload")
	}
	
	item := NewItem(updateItemRequest.Name)

	if err := s.database.UpdateItem(id, item); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, item)
}

func (s *APIServer) deleteItemHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, err.Error())
	}
	if err := s.database.DeleteItem(id); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func (s *APIServer) listItemHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, err.Error())
	}

	item, err := s.database.ListItemById(id)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err.Error())
	}

	return WriteJSON(w, http.StatusOK, item)
}

func WriteJSON(w http.ResponseWriter, statusCode int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(value) 
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"` 
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w,r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func getId(r *http.Request) (int, error) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idParam)
	}
	return id, nil
}