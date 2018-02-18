package hexmap

import "github.com/markllama/hexgame/types/hexvector"

type Terrain struct {
	Type string         `json:"type"`
	Name string         `json:"name,omitempty"`
	Locations []hexvector.Vector  `json:"locations,omitempty"`
	Regions []Region  `json:"regions,omitempty"`
}
