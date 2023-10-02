package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, Status int, v any) error {
	
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
type apiFunc func(http.ResponseWriter, *http.Request) error 

type ApiError struct {
	Error string
}
type makeHTTPHandleFunc(f apiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w, r); err != nil {
			WriterJSON(w, http.StatusBadRequest, ApiError{Error: err.Error})
		}
	}

}

type APIServer struct {
	ListenAddr string
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.Handelfunc("/account", makeHTTPHandleFunc(s.handelAccount)) 

	http.ListenAndServe(s.ListenAddr, router)
}

func NewAPIServer(mylisternAddr string) *APIServer {
	return &APIServer{
		ListenAddr: mylisternAddr,
	}
}

func (s *APIServer) handelAccount(w http.ResposeWriter, *http.Request ) error {
	if r.Method == "GET" {
		retrun s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		retrun s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		retrun s.handleDeleteAccount(w, r)
	}


	return fmt.Errorf("Method not allowed : %s", r.Method)
}

func (s *APIServer) handelGetAccount(w http.ResposeWriter, *http.Request ) error {
	return nil
}

func (s *APIServer) handelCreateAccount(w http.ResposeWriter, *http.Request ) error {
	return nil
}

func (s *APIServer) handelDeleteAccount(w http.ResposeWriter, *http.Request ) error {
	return nil
}

func (s *APIServer) handelTranser(w http.ResposeWriter, *http.Request ) error {
	return nil
}


