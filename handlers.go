package main

import (
	"encoding/json"
	"net/http"

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

	r.HandleFunc("/item", makeHTTPHandleFunc(s.handleItem))

	// endpoints, handler functions and HTTP methods
	// r.HandleFunc("/items", listItemsHandler).Methods("GET")
	// r.HandleFunc("/items", createItemHandler).Methods("POST")
	// r.HandleFunc("/items/{id}", updateItemHandler).Methods("PUT")
	// r.HandleFunc("/items/{id}", deleteItemHandler).Methods("DELETE")
	// r.HandleFunc("/items/{id}", listItemHandler).Methods("GET")

	http.ListenAndServe(s.listenAddress, r)
}

func (s *APIServer) handleItem(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.listItemHandler(w, r)
	}
	if r.Method == "POST" {
		return s.createItemHandler(w, r)
	}
	if r.Method == "DELETE" {
		return s.deleteItemHandler(w, r)
	}
	if r.Method == "PUT" {
		return s.updateItemHandler(w, r)
	}
	return nil
}

func (s *APIServer) listItemsHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) createItemHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) updateItemHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) deleteItemHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) listItemHandler(w http.ResponseWriter, r *http.Request) error {
	item := NewItem(1,"Laiane")

	return WriteJSON(w, http.StatusOK, item)
}

func WriteJSON(w http.ResponseWriter, statusCode int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(value) //items 
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w,r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}