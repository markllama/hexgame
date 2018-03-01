package api

import (
	"github.com/markllama/hexgame/types"
)

type Game struct {
	types.Game `json:",inline"`
	URL string `json:"url,omitempty" bson:"-"`
}

type GameRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
