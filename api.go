package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	listenAddress string
	Store Storage
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)

}

func NewAPIServer(listenAddress string) *APIServer {
  // storage, err := NewPostgresStore()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return &APIServer{
		listenAddress: listenAddress,
		// store: storage,
	}
}

func (s *APIServer) Run() {
	router := chi.NewRouter()

	router.Get("/budgets/{id}", makeHTTPHandleFunc(s.handleGetBudget))

	log.Println(("Starting server on " + s.listenAddress))

	http.ListenAndServe(s.listenAddress, router)
}

// * --- Budgets ---

func (s *APIServer) handleGetBudget(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}

	budget := NewBudget(id, "Test Budget", 0, 1)

	return WriteJSON(w, http.StatusOK, budget)
}

func (s *APIServer) handleCreateBudget(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteBudget(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// * --- Transactions ---
func (s *APIServer) handleCreateTransaction(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteTransaction(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// * --- Users ---
