package types

type Terrain struct {
	Type string         `json:"type"`
	Name string         `json:"name,omitempty"`
	Locations []Vector  `json:"locations"`
	// Regions []Region  `json:"regions,omitempty"`
}
