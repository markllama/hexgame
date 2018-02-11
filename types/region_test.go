package types

import (
	"testing"

	"fmt"
)

func TestRegion(t *testing.T) {
	r := CircularRegion{Vector{0, 0}, 1}

	fmt.Println(r.All())
	
}
