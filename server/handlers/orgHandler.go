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
	return utils.WriteJSON(w, http.StatusCreated, "Organization created successfully")
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
	return utils.WriteJSON(w, http.StatusOK, org)
}
