package hexmap

import (
	"testing"

	"github.com/markllama/hexgame/types/hexvector"
)


func TestRegionContains(t *testing.T) {

	m0 := Region{
		Origin: hexvector.ORIGIN,
		Size: hexvector.Vector{Hx: 15, Hy: 21 },
	}

	inside := []hexvector.Vector{
		hexvector.Vector{Hx: 1, Hy: 1},
		hexvector.Vector{Hx: 0, Hy: 20},
	}

	outside := []hexvector.Vector{
		hexvector.Vector{Hx: -1, Hy: 3},
		hexvector.Vector{Hx: 16, Hy: 10},
	}

	var actual bool

	actual = m0.Contains(hexvector.ORIGIN)
	if actual == false {
		t.Error("Origin is not inthe map")
	}

	for _, v := range inside {
		actual = m0.Contains(v)
		if actual == false {
			t.Error("hex", v, "is not contained and should be")
		}
	}

	for _, v := range outside {
		actual = m0.Contains(v)
		if actual == true {
			t.Error("hex", v, "is contained and should not be")
		}
	}
}


func TestRegionAll(t *testing.T) {

	r := Region{
		Origin: hexvector.ORIGIN,
		Size: hexvector.Vector{Hx: 4, Hy: 4},
	}

	hexes := r.All()

	if len(hexes) != 16 {
		t.Error("Should be 16 hexes")
	}

}
