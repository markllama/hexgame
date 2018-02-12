package types

import (
	"testing"

	"fmt"
)

func TestRegion(t *testing.T) {
	r := CircularRegion{Vector{0, 0}, 3}

	fmt.Println(r.All(nil))
	fmt.Println(r.All(&Map{Origin: Vector{3, 4}}))
	
}
