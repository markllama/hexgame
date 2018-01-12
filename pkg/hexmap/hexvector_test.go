package hexmap

import "testing"
import "fmt"
import "strings"

func TestConstructors(t *testing.T) {
	hv := HexVector{4, 1}

	if hv.Hx() != 4 {
		t.Error("hv.Hx(): expected 4, actual: " + fmt.Sprintf("d", hv.Hx()))
	}
	if hv.Hy() != 1 {
		t.Error("hv.Hy(): expected 1, actual: " + fmt.Sprintf("%d", hv.Hy()))
	}
	if hv.Hz() != -3 {
		t.Error("hv.Hz(): expected -3, actual: " + fmt.Sprintf("%d", hv.Hz()))
	}
}


func TestEqual(t *testing.T) {
	var o = HexVector{0, 0}

	if o.Equal(ORIGIN) != true {
		t.Error("HexVector.Equal() = origin mismatch")
	}

	o = HexVector{0, -1}
	if o.Equal(UNIT[0]) != true {
		t.Error("HexVector.Equal() = UNIT[0] mismatch")
	}
}

func TestAdd(t *testing.T) {
	var hv0 = HexVector{hx: 3, hy: 5}
	var hv1 = HexVector{hx: 4, hy: -6}

	var hv2 = hv0.Add(hv1)

	if hv2.Hx() != 7 {
		t.Error("HexVector.Add - Hx() expected: 7: actual: " + fmt.Sprintf("%d", hv2.Hx()))
	}
	if hv2.Hy() != -1 {
		t.Error("HexVector.Add - Hy() expected: -1 : actual: " + fmt.Sprintf("%d", hv2.Hy()))
	}
}

func TestSub(t *testing.T) {
	var hv0 = HexVector{hx: 3, hy: 5}
	var hv1 = HexVector{hx: 4, hy: -6}

	var hv2 = hv0.Sub(hv1)

	if hv2.Hx() != -1 {
		t.Error("HexVector.Sub - Hx() expected: -1: actual: " + fmt.Sprintf("%d", hv2.Hx()))
	}
	if hv2.Hy() != 11 {
		t.Error("HexVector.Sub - Hy() expected: 10 : actual: " + fmt.Sprintf("%d", hv2.Hy()))
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
		t.Error("HexVector.max3(): expected first: 3, actual: " + fmt.Sprintf("%d", a))	
	}

	a =  max3(2, 3, 1)
	if a != 3 {
		t.Error("HexVector.max3(): expected second: 3, actual: " + fmt.Sprintf("%d", a))	
	}


	a =  max3(2, 1, 3)
	if a != 3 {
		t.Error("HexVector.max3(): expected third: 3, actual: " + fmt.Sprintf("%d", a))	
	}

	a = max3(3, 3, 3)
	if a != 3 {
		t.Error("HexVector.max3(): expected all: 3, actual: " + fmt.Sprintf("%d", a))	
	}

	a = max3(-3, -3, -3)
	if a != -3 {
		t.Error("HexVector.max3(): expected all: -3, actual: " + fmt.Sprintf("%d", a))	
	}
}

func TestLength(t *testing.T) {
	var hx0 = HexVector{hx: 3, hy: 2}
	var hx1 = HexVector{hx: 4, hy: -3}
	
	var l = hx0.Length()
	if l != 3 {
		t.Error("HexVector.Length() expect: 3, actual: " + fmt.Sprintf("%d", l))
	}

	l = hx1.Length()
	if l != 7 {
		t.Error("HexVector.Length() expect: 7, actual: " + fmt.Sprintf("%d", l))
	}
}

func TestDistance(t* testing.T) {
 	var hv0 = HexVector{3, 4}
 	var hv1 = HexVector{6, -3}

 	var d = hv0.Distance(hv1)
	if d != 10 {
		t.Error("HexVector.Distance(): expected: 10, actual: " + fmt.Sprintf("%d", d))
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
		t.Error("HexVector.Hextant() ORIGIN, expected: 0, actual: " + fmt.Sprintf("0%03o", h))
	}

	for i, v := range UNIT {
		h = v.Hextant()
		if h != i % 6 {
			t.Error(fmt.Sprintf("HexVector.Hextant() UNIT[%d], expected: %d, actual %d",i, i % 6, h))
		}
	}

	var offaxis = []HexVector{
		HexVector{ 1,-1},
		HexVector{ 2, 1},
		HexVector{ 1, 2},
		HexVector{-1, 1},
		HexVector{-2,-1},
		HexVector{-1,-2},
	}

	for i, v := range offaxis {
		h = v.Hextant()
		if h != i {
			t.Error(fmt.Sprintf("HexVector.Hextant() UNIT[%d], expected: %d, actual %d",i, i, h))
		}
	}
}


func TestMarshal(t *testing.T) {
	hv0 := HexVector{-3, 5}

	hv_json, err := hv0.MarshalJSON()

	if err != nil {
		t.Error(fmt.Sprintf("HexVector.MarshalJSON(): error: %s", err))
	}

	hv_string := string(hv_json)
	if strings.Compare(hv_string, "{\"hx\":-3,\"hy\":5}") != 0 {
		t.Error(fmt.Sprintf("'%s'", hv_string))
	}
}



func TestUnmarshal(t *testing.T) {
	var hv0 HexVector
 	hv_string := "{\"hx\":-3,\"hy\":5}"

  err := hv0.UnmarshalJSON([]byte(hv_string))

	if err != nil {
		t.Error(fmt.Sprintf("HexVector.Unmarshal(): error: %s", err))
	}

	if hv0.Hx() != -3 {
		t.Error(fmt.Sprintf("HexVector.Unmarshal() hx:  expected 3, actual: %d", hv0.Hx()))
	}
}
