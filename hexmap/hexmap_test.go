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
