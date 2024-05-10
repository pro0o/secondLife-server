package handlers

import (
	"net/http"
	"secondLife/utils"

	"github.com/google/uuid"
)

func (h *APIServer) login(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email address not provided", http.StatusBadRequest)
		return nil
	}
	user, err := h.store.GetUserByEmail(email)
	if err != nil {
		return err
	}

	responseData := struct {
		UserID         uuid.UUID `json:"user_id"`
		ProfilePicture int       `json:"profile_picture"`
		UserName       string    `json:"user_name"`
	}{
		UserID:         user.UserID,
		ProfilePicture: user.ProfilePicture,
		UserName:       user.UserName,
	}

	if err := utils.WriteJSON(w, http.StatusOK, responseData); err != nil {
		return err
	}

	return nil
}
