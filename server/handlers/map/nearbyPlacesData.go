package osm

import (
	"encoding/json"
	"secondLife/types"
)

func GetNearbyData(longitude, latitude float64, object string) ([]types.Location, error) {
	var data types.JSONData
	if err := json.Unmarshal([]byte(JsonData), &data); err != nil {
		return nil, err
	}

	var nearbyLocations []types.Location

	switch object {
	case "clothes":
		nearbyLocations = append(nearbyLocations, data.Orphanages...)
		nearbyLocations = append(nearbyLocations, data.Ngos...)
	case "plastic":
		nearbyLocations = append(nearbyLocations, data.Plastic...)
		nearbyLocations = append(nearbyLocations, data.All...)
	// TODO: OTHER OBJECT NEARBY RETREIVING
	default:
		return nil, nil
	}

	var nearbyData []types.Location
	for _, loc := range nearbyLocations {
		distance := calculateDistance(latitude, longitude, loc.Latitude, loc.Longitude)
		if distance <= 5.0 { // within 5 kilometers
			nearbyData = append(nearbyData, loc)
		}
	}

	return nearbyData, nil
}
