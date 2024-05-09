package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"secondLife/types"
	"secondLife/utils"

	"github.com/google/uuid"
)

func (h *APIServer) rewardPointsHandler(w http.ResponseWriter, r *http.Request) error {
	switch {
	case r.Method == http.MethodPost:
		return h.updateRewardPointsHandler(w, r)
	case r.Method == http.MethodGet:
		return h.getRewardPointsHandler(w, r)
	default:
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: "invalid request method"})
	}
}

func (h *APIServer) updateRewardPointsHandler(w http.ResponseWriter, r *http.Request) error {
	var rewardPoints types.RewardPoints
	if err := json.NewDecoder(r.Body).Decode(&rewardPoints); err != nil {
		return err
	}

	if rewardPoints.UserID == uuid.Nil {
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: "userID cannot be empty"})
	}

	err := h.store.UpdateUserPoints(rewardPoints.UserID, rewardPoints.Points)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true})
}

func (h *APIServer) getRewardPointsHandler(w http.ResponseWriter, r *http.Request) error {
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: "userID parameter is missing"})
	}

	parsedUserID, err := uuid.Parse(userID)
	log.Println("userID is:", parsedUserID)
	if err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: "invalid userID format"})
	}
	points, err := h.store.GetUserPointsByID(parsedUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.WriteJSON(w, http.StatusNotFound, ApiError{Error: "user not found"})
		}
		return err
	}
	response := struct {
		Points int `json:"points"`
	}{
		Points: points,
	}
	return utils.WriteJSON(w, http.StatusOK, response)
}
