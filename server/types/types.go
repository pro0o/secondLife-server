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
