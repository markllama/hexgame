package hexmap

import (
	"github.com/markllama/hexgame/types/hexvector"
)

type Map struct {
	Name string `json:"name"`
	Copyright string `json:"copyright"`
	Shape string `json:"shape"`
	Size hexvector.Vector `json:"size" bson:"size"`
	Origin hexvector.Vector `json:"origin" bson:"origin"`
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


func ybias(hx int) (int) {
	// 0 -> 0
	// 1 -> 0
	// 2 -> 1
	// 3 -> 1

	if hx >= 0  { return hx / 2 }

	// -1 -> -1
	// -2 -> -1
	// -3 -> -2
	// -3 -> -2
	return -((-hx + 1) / 2)
}

func (m *Map) Contains(v hexvector.Vector) (bool) {

	// Test if it's off the left or right side
	if v.Hx < m.Origin.Hx || v.Hx >= m.Origin.Hx + m.Size.Hx {
		return false
	}
	
	bias := ybias(v.Hx)
	if v.Hy < m.Origin.Hy + bias || v.Hy >= m.Origin.Hy + m.Size.Hy + bias {
		return false
	}
	
	return true
}
