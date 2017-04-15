package hexmap

type HexVector struct {
	hx int
	hy int
}


// Accessor functions
func (hv HexVector) Hx() (int) {
	return hv.hx
}

func (hv HexVector) Hy() (int) {
	return hv.hy
}

func (hv HexVector ) Hz() (int) {
	return hv.hy - hv.hx
}

func (hv0 HexVector) Add(hv1 HexVector) (HexVector) {
	return HexVector{hv0.Hx() + hv1.Hx(), hv0.Hy() + hv1.Hy()}
}

func (hv0 HexVector) Sub(hv1 HexVector) (HexVector) {
	return HexVector{hv0.Hx() - hv1.Hx(), hv0.Hy() - hv1.Hy()}
}

func abs(i int) (a int) {
	if i == 0 {
		a = 0
	} else {
		a = i * (i / i)
	}
	return 
}

func max3(i, j, k int) (a int) {
	if i < j {
		a = j
	} else {
		a = i
	}
	if a < k {
		a = k
	}
	return
}

func (hv HexVector) Length() (l int) {
	return max3(abs(hv.Hx()), abs(hv.Hy()), abs(hv.Hz()))
}

func (hv0 HexVector) Distance(hv1 HexVector) (int) {
	return  abs(hv0.Sub(hv1).Length())
}
