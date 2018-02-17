package types

//
// This class represents a set of hexes on a map.
//


type Region struct {
	Origin Vector `bson:"origin"`
	Size Vector `bson:"size"`
}


func (r Region) Contains(v Vector, m *Map) {
	
	
	
}

/*
type Region interface {
	All() []Vector
	Contains(Vector) bool
}

type CircularRegion struct {
	Center Vector `json:"center"`
	Radius int `json:"radius"`
}
	
func (region CircularRegion) All(m *Map) ([]Vector) {

	var hexes []Vector
	var origin Vector
	//var size Vector
	
	if m != nil {
		origin = m.Origin
	} else {
		origin = ORIGIN
	}

 	// start at -x, scan to +x, filling in all y
 	for hx := -(region.Radius - 1) ; hx <= region.Radius - 1 ; hx++ {
		var hy_low, hy_high int
		if hx <= 0 {
			hy_low = -(region.Radius -1)
			hy_high = (region.Radius -1) + hx
		} else {
			hy_low = -(region.Radius - 1) + hx
			hy_high = region.Radius - 1
		}
 		for hy := hy_low ; hy <= hy_high  ; hy++ {
			hex := Vector{hx + origin.Hx, hy + origin.Hy}
			if m == nil || m.Contains(hex) {
				hexes = append(hexes, hex)
			}
		}
 	}
	
	return hexes
}

func (r CircularRegion) Contains(location Vector, m *Map) (bool) {
	// if the location is within Radius hexes of Center

	return m.Contains(location) && location.Distance(r.Center) <= r.Radius
}

*/

