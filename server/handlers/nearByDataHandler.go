package handlers

import (
	"encoding/json"
	"net/http"
	osm "secondLife/handlers/map"
	"secondLife/types"
)

func (h *APIServer) handleNearByData(w http.ResponseWriter, r *http.Request) error {
	var data types.DataFromClientMap
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}

	nearbyData, err := osm.GetNearbyData(data.Longitude, data.Latitude, data.Object)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(nearbyData); err != nil {
		return err
	}

	return nil
}
