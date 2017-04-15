package hexmap

import "testing"
import "fmt"

func TestConstructors(t *testing.T) {
	var hv = HexVector{hx: 4, hy: 1}

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
