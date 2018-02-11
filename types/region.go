package types

/*
This class represents a set of hexes on a map.
*/

type Region interface {
	All() []Vector
	Contains(Vector) bool
}


type CircularRegion struct {
	Center Vector `json:"center"`
	Radius int `json:"radius"`
}
	
func (r CircularRegion) All() ([]Vector) {

	var v []Vector

 	// start at -x, scan to +x, filling in all y
 	for hx := -(r.Radius - 1) ; hx <= r.Radius - 1 ; hx++ {
		var hy_low, hy_high int
		if hx <= 0 {
			hy_low = -(r.Radius -1)
			hy_high = (r.Radius -1) + hx
		} else {
			hy_low = -(r.Radius - 1) + hx
			hy_high = r.Radius - 1
		}
 		for hy := hy_low ; hy <= hy_high  ; hy++ {
			v = append(v, Vector{hx, hy})
		}
 	}
	
	return v
}

func (r CircularRegion) Contains(location Vector) (bool) {
	// if the location is within Radius hexes of Center
	return location.Distance(r.Center) <= r.Radius
}

type LineRegion struct {
	// define a line and which side of the line is "inside"
	Axis Axis
	Side bool  // false = Less, true = More
}
