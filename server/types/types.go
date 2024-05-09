package types

import "github.com/google/uuid"

type Video struct {
	Image  string `json:"image"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author string `json:"author"`
}

type OverpassResponse struct {
	Lat  float64           `json:"lat"`
	Lon  float64           `json:"lon"`
	Tags map[string]string `json:"tags"`
	Type string            `json:"type"`
}

type User struct {
	UserID         uuid.UUID `json:"user_id"`
	Email          string    `json:"email"`
	ProfilePicture int       `json:"profile_picture"`
	UserName       string    `json:"user_name"`
}

type Org struct {
	UserID      uuid.UUID `json:"user_id"`
	OrgName     string    `json:"org_name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
}

type RewardPoints struct {
	UserID uuid.UUID `json:"userID"`
	Points int       `json:"points"`
}

type DataFromClientMap struct {
	Object    string  `json:"object"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Location struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Location  string  `json:"location"`
}

type JSONData struct {
	Orphanages    []Location `json:"orphanages"`
	Food          []Location `json:"food"`
	RecycleCentre []Location `json:"recycleCentre"`
	Ngos          []Location `json:"ngos"`
	Clothes       []Location `json:"clothes"`
	EWaste        []Location `json:"E-Waste"`
	Plastic       []Location `json:"plastic"`
	Glass         []Location `json:"glass"`
	Paper         []Location `json:"paper"`
	All           []Location `json:"all"`
}
