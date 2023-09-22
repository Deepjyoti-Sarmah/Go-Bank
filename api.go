package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
  w.WriteHeader(status) 
  w.Header().Set("Content Type", "application/json")
  return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error 

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {
    if err := f(w, r); err != nil {
      //handle the error 
    }
  }
}

type APIServer struct {
  listenAddr string
}

func newAPIServer(listenAddr string) *APIServer {
  return &APIServer{
    listenAddr: listenAddr,
  }
}

func (s *APIServer) Run() {
  router := mux.NewRouter()

  router HandleFunc("/account", s.handleAccount)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error  {
  return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error  {
  return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error  {
  return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error  {
  return nil
}

func (s *APIServer) hancleTransfer(w http.ResponseWriter, r *http.Request) error  {
  return nil
}
