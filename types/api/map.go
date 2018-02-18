package api

import (
	"github.com/markllama/hexgame/types/hexmap"
)

type MapRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Map struct {
	Game string `json:"game,omitempty"`
	hexmap.Map `json:",inline"`
	URL string `json:"url,omitempty"`
}
