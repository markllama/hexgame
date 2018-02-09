package types

type Map struct {
	Name string `json:"name"`
	Copyright string `json:"copyright"`
	Shape string `json:"shape"`
	Size Vector `json:"size" bson:"size"`
	Origin Vector `json:"origin" bson:"origin"`
	Terrains []Terrain `json:"terrains"`
	Tokens []Token `json:"tokens,omitempty"`
}

func (m0 *Map) Copy(m1 Map) {
	m0.Name = m1.Name
	m0.Copyright = m1.Copyright
	m0.Shape = m1.Shape
	
	m0.Size.Hx = m1.Size.Hx
	m0.Size.Hy = m1.Size.Hy
//	m0.Size.Copy(m1.Size)
//	m0.Origin.Copy(m1.Origin)

	// copy all the terrains

	// copy all the tokens
}
