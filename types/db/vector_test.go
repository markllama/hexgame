package hexmap

import "testing"
import "fmt"
import "encoding/json"
import "strings"

func TestConstructors(t *testing.T) {
	v := Vector{4, 1}

	if v.Hx != 4 {
		t.Error("v.Hx: expected 4, actual: " + fmt.Sprintf("d", v.Hx))
	}
	if v.Hy != 1 {
		t.Error("v.Hy: expected 1, actual: " + fmt.Sprintf("%d", v.Hy))
	}
	if v.Hz() != -3 {
		t.Error("v.Hz(): expected -3, actual: " + fmt.Sprintf("%d", v.Hz()))
	}
}


func TestEqual(t *testing.T) {
	var o = Vector{0, 0}

	if o.Equal(ORIGIN) != true {
		t.Error("Vector.Equal() = origin mismatch")
	}

	o = Vector{0, -1}
	if o.Equal(UNIT[0]) != true {
		t.Error("Vector.Equal() = UNIT[0] mismatch")
	}
}

func TestAdd(t *testing.T) {
	var v0 = Vector{Hx: 3, Hy: 5}
	var v1 = Vector{Hx: 4, Hy: -6}

	var v2 = v0.Add(v1)

	if v2.Hx != 7 {
		t.Error("Vector.Add - Hx expected: 7: actual: " + fmt.Sprintf("%d", v2.Hx))
	}
	if v2.Hy != -1 {
		t.Error("Vector.Add - Hy expected: -1 : actual: " + fmt.Sprintf("%d", v2.Hy))
	}
}

func TestSub(t *testing.T) {
	var v0 = Vector{Hx: 3, Hy: 5}
	var v1 = Vector{Hx: 4, Hy: -6}

	var v2 = v0.Sub(v1)

	if v2.Hx != -1 {
		t.Error("Vector.Sub - Hx expected: -1: actual: " + fmt.Sprintf("%d", v2.Hx))
	}
	if v2.Hy != 11 {
		t.Error("Vector.Sub - Hy expected: 10 : actual: " + fmt.Sprintf("%d", v2.Hy))
	}
}


func TestAbs(t *testing.T) {
	if abs(4) != 4 {
		t.Error("Hexvector.abs(): expected: 4, actual: " + fmt.Sprintf("%d", abs(4)))
	}

	if abs(-4) != 4 {
		t.Error("Hexvector.abs(): expected: 4, actual: " + fmt.Sprintf("%d", abs(-4)))
	}
}

func TestMax3(t *testing.T) {
	var a =  max3(3, 2, 1)
	if a != 3 {
		t.Error("Vector.max3(): expected first: 3, actual: " + fmt.Sprintf("%d", a))	
	}

	a =  max3(2, 3, 1)
	if a != 3 {
		t.Error("Vector.max3(): expected second: 3, actual: " + fmt.Sprintf("%d", a))	
	}


	a =  max3(2, 1, 3)
	if a != 3 {
		t.Error("Vector.max3(): expected third: 3, actual: " + fmt.Sprintf("%d", a))	
	}

	a = max3(3, 3, 3)
	if a != 3 {
		t.Error("Vector.max3(): expected all: 3, actual: " + fmt.Sprintf("%d", a))	
	}

	a = max3(-3, -3, -3)
	if a != -3 {
		t.Error("Vector.max3(): expected all: -3, actual: " + fmt.Sprintf("%d", a))	
	}
}

func TestLength(t *testing.T) {
	var Hx0 = Vector{Hx: 3, Hy: 2}
	var Hx1 = Vector{Hx: 4, Hy: -3}
	
	var l = Hx0.Length()
	if l != 3 {
		t.Error("Vector.Length() expect: 3, actual: " + fmt.Sprintf("%d", l))
	}

	l = Hx1.Length()
	if l != 7 {
		t.Error("Vector.Length() expect: 7, actual: " + fmt.Sprintf("%d", l))
	}
}

func TestDistance(t* testing.T) {
 	var v0 = Vector{3, 4}
 	var v1 = Vector{6, -3}

 	var d = v0.Distance(v1)
	if d != 10 {
		t.Error("Vector.Distance(): expected: 10, actual: " + fmt.Sprintf("%d", d))
	}
}


func TestHextant(t* testing.T) {

	if hextant[0111] != 0 {
		t.Error("hextant[0111] expect: 0, actual: " + fmt.Sprintf("%d", hextant[0111]))
	}

	if hextant[0100] != 0 {
		t.Error("hextant[0100] expect: 0, actual: " + fmt.Sprintf("%d", hextant[0100]))
	}

	if hextant[0210] != 1 {
		t.Error("hextant[0210] expect: 1, actual: " + fmt.Sprintf("%d", hextant[0210]))
	}

	if hextant[0221] != 2 {
		t.Error("hextant[0221] expect: 2, actual: " + fmt.Sprintf("%d", hextant[0221]))
	}

	if hextant[0200] != 0 {
		t.Error("hextant[0200] expect: 0, actual: " + fmt.Sprintf("%d", hextant[0200]))
	}

	var h int = ORIGIN.Hextant()
	if h != 0 {
		t.Error("Vector.Hextant() ORIGIN, expected: 0, actual: " + fmt.Sprintf("0%03o", h))
	}

	for i, v := range UNIT {
		h = v.Hextant()
		if h != i % 6 {
			t.Error(fmt.Sprintf("Vector.Hextant() UNIT[%d], expected: %d, actual %d",i, i % 6, h))
		}
	}

	var offaxis = []Vector{
		Vector{ 1,-1},
		Vector{ 2, 1},
		Vector{ 1, 2},
		Vector{-1, 1},
		Vector{-2,-1},
		Vector{-1,-2},
	}

	for i, v := range offaxis {
		h = v.Hextant()
		if h != i {
			t.Error(fmt.Sprintf("Vector.Hextant() UNIT[%d], expected: %d, actual %d",i, i, h))
		}
	}
}


func TestMarshal(t *testing.T) {
	v0 := Vector{Hx: -3, Hy: 5}

	v_json, err := json.Marshal(v0)
	
// 	v_json, err := v0.MarshalJSON()

 	if err != nil {
 		t.Error(fmt.Sprintf("Vector.MarshalJSON(): error: %s", err))
	}

 	v_string := string(v_json)
 	if strings.Compare(v_string, "{\"hx\":-3,\"hy\":5}") != 0 {
 		t.Error(fmt.Sprintf("expected {hx:-3,hy:5}: actual: '%s'", v_json))
	}
}



func TestUnmarshal(t *testing.T) {
 	var v0 Vector
	v_string := "{\"hx\":-3,\"hy\":5}"

	err := json.Unmarshal([]byte(v_string), &v0)

 	if err != nil {
 		t.Error(fmt.Sprintf("Vector.Unmarshal(): error: %s", err))
 	}

 	if v0.Hx != -3 {
 		t.Error(fmt.Sprintf("Vector.Unmarshal() Hx:  expected -3, actual: %d", v0.Hx))
 	}
}
