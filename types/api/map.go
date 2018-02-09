package api

import (
	"github.com/markllama/hexgame/types"
)

type MapRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Map struct {
	Game string `json:"game,omitempty"`
	types.Map `json:",inline"`
	URL string `json:"url,omitempty"`
}
