package hexmap

import "testing"
import "strings"

func TestTerrainConstructor(t *testing.T) {
	t0 := Terrain{"clear", nil, nil}

	if strings.Compare("clear", t0.Name) != 0 {
		t.Error("Name = " + t0.Name)
	}
}
