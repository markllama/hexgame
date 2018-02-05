package api

import (
	"github.com/markllama/hexgame/types"
)

type MapRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Map struct {
	Name string `json:"name"`
	Game string `json:"game"`
	Copyright string `json:"copyright"`
	Shape string `json:"shape"`
	Size types.Vector `json:"size"`
	Origin types.Vector `json:"origin"`
	Terrains []types.Terrain `json:"terrains"`
	Tokens []types.Token `json:"tokens,omitempty"`
	URL string `json:"url,omitempty"`
}
