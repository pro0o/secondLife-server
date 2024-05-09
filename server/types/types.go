package types

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
