package api

import (
	"github.com/markllama/hexgame/types/hexvector"
)

type Terrain struct {
	Name      string              `json:"name"`
	Locations []hexvector.Vector  `json:"locations"`
}
