package handlers

import (
	"encoding/json"
	"net/http"
	"secondLife/types"
	"secondLife/utils"
)

func (h *APIServer) orgHandler(w http.ResponseWriter, r *http.Request) error {
	switch {
	case r.Method == http.MethodPost:
		return h.createOrgHandler(w, r)
	case r.Method == http.MethodGet:
		return h.getOrgHandler(w, r)
	default:
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: "invalid request method"})
	}
}

func (h *APIServer) createOrgHandler(w http.ResponseWriter, r *http.Request) error {
	var orgData types.Org
	if err := json.NewDecoder(r.Body).Decode(&orgData); err != nil {
		return err
	}

	if err := h.store.CreateOrg(&orgData); err != nil {
		return err
	}
	userName, err := h.store.GetUserName(orgData.UserID, orgData.OrgName)
	if err != nil {
		return err
	}
	// Return created organization data
	responseData := struct {
		UserName    string `json:"user_name"`
		OrgName     string `json:"org_name"`
		Location    string `json:"location"`
		Description string `json:"description"`
	}{
		UserName:    userName,
		OrgName:     orgData.OrgName,
		Location:    orgData.Location,
		Description: orgData.Description,
	}

	return utils.WriteJSON(w, http.StatusCreated, responseData)
}

// /orgs?orgName=example_org_name
func (h *APIServer) getOrgHandler(w http.ResponseWriter, r *http.Request) error {
	orgName := r.URL.Query().Get("org_name")
	if orgName == "" {
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: "org_name parameter is missing"})
	}

	org, err := h.store.GetOrgByName(orgName)
	if err != nil {
		return err
	}
	userName, err := h.store.GetUserName(org.UserID, org.OrgName)
	if err != nil {
		return err
	}
	responseData := struct {
		UserName    string `json:"user_name"`
		OrgName     string `json:"org_name"`
		Location    string `json:"location"`
		Description string `json:"description"`
	}{
		UserName:    userName,
		OrgName:     org.OrgName,
		Location:    org.Location,
		Description: org.Description,
	}
	return utils.WriteJSON(w, http.StatusOK, responseData)
}
