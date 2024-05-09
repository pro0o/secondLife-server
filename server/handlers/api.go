package handlers

import (
	"log"
	"net/http"
	"secondLife/storage"
	"secondLife/utils"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      *storage.PostgresStore
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string, store *storage.PostgresStore) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
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
	router.HandleFunc("/signUp", makeHTTPHandleFunc(h.signUp))
	router.HandleFunc("/login", makeHTTPHandleFunc(h.handleSendVerificationCode))
	router.HandleFunc("/org", makeHTTPHandleFunc(h.orgHandler))
	router.HandleFunc("/rewardPoints", makeHTTPHandleFunc(h.rewardPointsHandler))
	router.HandleFunc("/recyclingData/videos", makeHTTPHandleFunc(h.handleYoutubeSuggestion))
	router.HandleFunc("/recyclingData/maps", makeHTTPHandleFunc(h.handleNearByData))

	log.Print("Server running on port", h.listenAddr)
	http.ListenAndServe(":8081", router)
}
