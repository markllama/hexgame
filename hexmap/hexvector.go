package hexmap

type HexVector struct {
	hx int
	hy int
}

var ORIGIN = HexVector{0, 0}
var UNIT = []HexVector{
	HexVector{ 0, -1},
	HexVector{ 1,  0},
	HexVector{ 1,  1},
	HexVector{ 0,  1},
	HexVector{-1,  0},
	HexVector{-1, -1},
	HexVector{ 0, -1},
}

/*
# -1 = 0
#  0 = 1
#  1 = 2

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

// Accessor functions
func (hv HexVector) Hx() (int) {
	return hv.hx
}

func (hv HexVector) Hy() (int) {
	return hv.hy
}

func (hv HexVector) Hz() (int) {
	return hv.hy - hv.hx
}

func (hv0 HexVector) Equal(hv1 HexVector) (bool) {
	return hv0.Hx() == hv1.Hx() && hv0.Hy() == hv1.Hy()
}

func (hv0 HexVector) Add(hv1 HexVector) (HexVector) {
	return HexVector{hv0.Hx() + hv1.Hx(), hv0.Hy() + hv1.Hy()}
}

func (hv0 HexVector) Sub(hv1 HexVector) (HexVector) {
	return HexVector{hv0.Hx() - hv1.Hx(), hv0.Hy() - hv1.Hy()}
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

func (hv HexVector) Length() (l int) {
	return max3(abs(hv.Hx()), abs(hv.Hy()), abs(hv.Hz()))
}

func (hv0 HexVector) Distance(hv1 HexVector) (int) {
	return  abs(hv0.Sub(hv1).Length())
}

// hx / abs(hx)

func (hv HexVector) Hextant() (int) {

	var m, rx, ry, rz int
	if hv.Hx() == 0 {	rx = 1 } else { rx = hv.Hx() / abs(hv.Hx()) + 1}		
	if hv.Hy() == 0 {	ry = 1 } else { ry = hv.Hy() / abs(hv.Hy()) + 1}
	if hv.Hz() == 0 {	rz = 1 } else { rz = hv.Hz() / abs(hv.Hz()) + 1}
	m = (rx << 6) + (ry << 3) + rz
	//return m

	return hextant[m]
}
