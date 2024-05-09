package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIServer struct{}

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer() *APIServer {
	return &APIServer{}
}

func (h *APIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/recognize":
		if r.Method == http.MethodPost {
			h.imageRecognize(w, r)
			return
		}
	default:
		http.NotFound(w, r)
		return
	}
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (h *APIServer) Run() {
	log.Println("Server running...")
	http.ListenAndServe(":8080", h)
}
