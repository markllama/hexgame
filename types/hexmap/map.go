package hexmap

import (
	"gopkg.in/mgo.v2"
)

type Map struct {
	Name string `json:"name"`
	Game string `json:"game"`
	Size Vector `json:"size"`
	Origin Vector `json:"origin"`
	// Terrains []Terrain `json:"terrains,omitempty"`
	// Tokens []Token `json:"tokens,omitempty"`
}
