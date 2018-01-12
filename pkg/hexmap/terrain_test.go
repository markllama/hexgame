package hexmap

import (
	"testing"
	"strings"
	"fmt"
)

func TestTerrainConstructor(t *testing.T) {
	t0 := Terrain{Name: "clear"}

	if strings.Compare("clear", t0.Name) != 0 {
		t.Error("Name = " + t0.Name)
	}
}

func TestTerrainLocations(t *testing.T) {
	loclist := []HexVector{
		{3, 4}, {4, 5}, {-2, -6},
	}
	
	t0 := Terrain{Name: "testloc", Locations: loclist}

	if len(t0.Locations) != 3 {
		t.Error(fmt.Sprintf("Terrain Locations - incorrect length expected 3, actual %d", len(t0.Locations)))
	}

}
