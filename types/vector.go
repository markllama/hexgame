
package types

//import (
//	"encoding/json"
//}

type Vector struct {
	Hx int `json:"hx" bson:"hx"`
	Hy int `json:"hy" bson:"hy"`
}

var ORIGIN = Vector{0, 0}
var UNIT = []Vector{
	Vector{ 0, -1},
	Vector{ 1,  0},
	Vector{ 1,  1},
	Vector{ 0,  1},
	Vector{-1,  0},
	Vector{-1, -1},
	Vector{ 0, -1},
}

/*
# on axis
#  0-1-1 0  100  0b010000
#  1 0-1 1  210  0b100100
#  1 1 0 2  221  0b101001
#  0 1 1 3  122  0b011010
# -1 0 1 4  012  0b000110
# -1-1 0 5  001  0b000001

# off axis
#  1-1-1  0 200  0b100000
#  1 1-1  1 220  0b101000
#  1 1 1  2 222  0b101010
# -1 1 1  3 022  0b001010
# -1-1 1  4 002  0b000010
# -1-1-1  5 000  0b000000
*/

var hextant = map[int]int{
	// origin
	0111: 0,
	
	// on axis
	0100: 0,
	0210: 1,
	0221: 2,
	0122: 3,
	0012: 4,
	0001: 5,

	// off axis
	0200: 0,
	0220: 1,
	0222: 2,
	0022: 3,
	0002: 4,
	0000: 5,
}

func (v Vector) Hz() (int) {
	return v.Hy - v.Hx
}

func (v0 *Vector) Copy(v1 Vector) {
	v0.Hx = v1.Hx
	v0.Hy = v1.Hy
}

func (v0 Vector) Equal(v1 Vector) (bool) {
	return v0.Hx == v1.Hx && v0.Hy == v1.Hy
}


func (v0 Vector) Add(v1 Vector) (Vector) {
	return Vector{v0.Hx + v1.Hx, v0.Hy + v1.Hy}
}

func (v0 Vector) Sub(v1 Vector) (Vector) {
	return Vector{v0.Hx - v1.Hx, v0.Hy - v1.Hy}
}

func abs(i int) (a int) {
	if i < 0 {	a = -i	} else { a = i	}
	return
}


func max3(i, j, k int) (a int) {
	if i < j { a = j} else { a = i }
	if a < k { a = k }
	return
}

func (v Vector) Length() (l int) {
	return max3(abs(v.Hx), abs(v.Hy), abs(v.Hz()))
}

func (v0 Vector) Distance(v1 Vector) (int) {
	return  abs(v0.Sub(v1).Length())
}


func (v Vector) Hextant() (int) {
	var m, rx, ry, rz int
	
	if v.Hx == 0 {	rx = 1 } else { rx = v.Hx / abs(v.Hx) + 1}
	if v.Hy == 0 {	ry = 1 } else { ry = v.Hy / abs(v.Hy) + 1}
	if v.Hz() == 0 {	rz = 1 } else { rz = v.Hz() / abs(v.Hz()) + 1}
	
	m = (rx << 6) + (ry << 3) + rz
	
	return hextant[m]
}

/*
func (v *Vector) UnmarshalJSON(b []byte) error {

}


func (v Vector) MarshalJSON () (b []byte, err error) {
	
}
*/
