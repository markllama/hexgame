package hexmap

type Terrain struct {
	Type string         `json:"type"`
	Locations []Vector  `json:"locations"`
}
