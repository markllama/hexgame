package api

import (
	"github.com/markllama/hexgame/types"
)

type MapRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Map struct {
	types.Map
	URL string `json:"url,omitempty"`
}
