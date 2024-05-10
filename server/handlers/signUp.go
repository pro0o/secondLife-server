package handlers

import (
	"encoding/json"
	"net/http"
	"secondLife/types"
	"secondLife/utils"

	"github.com/google/uuid"
)

func (h *APIServer) signUp(w http.ResponseWriter, r *http.Request) error {
	var newUser types.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return err
	}

	if h.store.CheckEmailExists(newUser.Email) {
		return utils.WriteJSON(w, http.StatusForbidden, "user already exists")
	}

	if err := h.store.CreateUser(&newUser); err != nil {
		return err
	}
	userID, err := h.store.GetUserUUID(newUser.Email)
	if err != nil {
		return err
	}

	responseData := struct {
		UserID         uuid.UUID `json:"user_id"`
		ProfilePicture int       `json:"profile_picture"`
		UserName       string    `json:"user_name"`
	}{
		UserID:         userID,
		ProfilePicture: newUser.ProfilePicture,
		UserName:       newUser.UserName,
	}

	if err := utils.WriteJSON(w, http.StatusOK, responseData); err != nil {
		return err
	}

	return nil
}
