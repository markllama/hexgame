package hexmap

import (
	"testing"
	"fmt"

	"github.com/markllama/hexgame/types/hexvector"
)



func TestYBias(t *testing.T) {

	values := [][]int{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 3},

		{-1, -1},
		{-2, -1},
		{-3, -2},
		{-4, -2},
	}
	
	var actual int

	for _, v := range values {
		actual = ybias(v[0])
		if actual != v[1] {
			t.Error(fmt.Sprintf("ybias fail: hx: %d, expected: %d, actual: %d",
				v[0], v[1], actual))
		}
	}
}


func TestMapContains(t *testing.T) {

	m0 := Map{
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
