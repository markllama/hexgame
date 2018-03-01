package types

import (
	"testing"
)

func TestYbias(t *testing.T) {

	input  :=    []int{0, 1, 2, 3, 4, 5, 6, -1, -2, -3, -4, -5, -6}
	expected :=  []int{0, 0, 1, 1, 2, 2, 3, -1, -1, -2, -2, -3, -3}

	m := Map{}

	for i := range input {
		actual := m.Ybias(input[i])
		if expected[i] != actual {
			t.Fatalf("Expected: %d, Actual: %d", expected[i], actual)
		}
	}
}

func TestContains(t *testing.T) {
	m := Map{Origin: ORIGIN, Size: Vector{15, 23}}

	inputs := []Vector{
		Vector{0,  0},
		Vector{-1, 0},
	}

	expected :=  []bool{
		true,
		false,
	}

	for i := range inputs {
		actual := m.Contains(inputs[i])
		if actual != expected[i] {
			t.Fatalf("expected: %s, actual: %s", expected, actual)
		}
	}
}
