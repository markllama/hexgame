package hexmap

import "testing";
import "fmt";

func TestHexMapConstructor(t *testing.T) {
	hm0 := HexMap{Name: "testmap", Size: &HexVector{3, 5}, Origin: &ORIGIN}

	if hm0.Name != "testmap" {
		t.Error("HexMap.Name - expected: 'testmap', actual: '" +  hm0.Name + "'")
	}

	if ! hm0.Size.Equal(HexVector{3, 5}) {
		t.Error(fmt.Sprintf("HexMap.Size - expected: {3, 5}, actual: {%d, %d}", hm0.Size.Hx(), hm0.Size.Hy()))
	}

	if ! hm0.Origin.Equal(ORIGIN) {
		t.Error(fmt.Sprintf("HexMap.Size - expected: {0, 0}, actual: {%d, %d}", hm0.Origin.Hx(), hm0.Origin.Hy()))
	}

}
