package handlers

import (
	"log"
	"net/http"
	"secondLife/utils"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func (h *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/recyclingData/videos-and-maps", makeHTTPHandleFunc(h.handleRecyclingData))

	log.Print("Server running on port", h.listenAddr)
	http.ListenAndServe(":8081", router)
}
