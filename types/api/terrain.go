package api

import (
	"github.com/markllama/hexgame/types"
)

type Terrain struct {
	Name string         `json:"name"`
	Locations []types.Vector  `json:"locations"`
}
