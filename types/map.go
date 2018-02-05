package types

type Map struct {
	Name string `json:"name"`
	Game string `json:"game"`
	Copyright string `json:"copyright"`
	Shape string `json:"shape"`
	Size Vector `json:"size"`
	Origin Vector `json:"origin"`
	Terrains []Terrain `json:"terrains"`
	Tokens []Token `json:"tokens,omitempty"`
}
