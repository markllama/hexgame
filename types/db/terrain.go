package hexmap

type Terrain struct {
	Name string         `json:"name"`
	Locations []Vector  `json:"locations"`
}
