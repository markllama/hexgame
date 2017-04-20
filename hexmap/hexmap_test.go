package hexmap

import (
	"testing"
	"fmt"
	"encoding/json"
	"strings"
)

func TestHexMapConstructor(t *testing.T) {
	hm0 := HexMap{Name: "testmap", Size: HexVector{3, 5}, Origin: ORIGIN}

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

func TestHexMapMarshal(t *testing.T) {

	hm := HexMap{Name: "TestMap", Size: HexVector{22, 14}, Origin: ORIGIN}

	jhm, err := json.Marshal(hm)

	if err != nil {
		t.Error(fmt.Sprintf("HexMap.Marshal() - error: %s", err))
	}

	if strings.Compare(string(jhm), "{\"name\":\"TestMap\",\"size\":{\"hx\":22,\"hy\":14},\"origin\":{\"hx\":0,\"hy\":0}}") != 0 { 
		t.Error(fmt.Sprintf("HexMap.Marshal(), - result: %s", jhm))
	}
}


// test containment and bias

func TestYbias(t *testing.T) {
	hm0 := HexMap{"ybiastest", HexVector{22,14}, ORIGIN}

	values := [][]int{{-3,-2}, {-2,-1}, {-1,-1}, {0,0}, {1,0}, {2,1}, {3,1}, {4,2}}

	for _, pair := range values {
		if hm0.ybias(pair[0]) != pair[1] {
			t.Error(fmt.Sprintf("HexMap.ybias() - %d -> expect: %d, actual: %d", pair[0], pair[1], hm0.ybias(pair[0])))

		}
	}	
}


//func TestContains(t *testing.T) {
//	
//}
