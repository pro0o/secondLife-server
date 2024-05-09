package handlers

import (
	"encoding/json"
	"net/http"

	"secondLife/handlers/youtube"
	"secondLife/parser"
	"secondLife/utils"
)

type JSONObject struct {
	ObjectDetected string `json:"objectDetected"`
}

func (h *APIServer) handleYoutubeSuggestion(w http.ResponseWriter, r *http.Request) error {
	var obj JSONObject
	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid JSON request body"})
	}

	objectString := parser.ReplaceSpacesWithPlus(obj.ObjectDetected)

	youtubeData, err := youtube.GetYoutubeData(objectString)
	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Error fetching YouTube data"})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(youtubeData)
	if err != nil {
		return err
	}

	return nil
}
