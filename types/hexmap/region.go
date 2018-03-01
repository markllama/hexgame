package hexmap

import "github.com/markllama/hexgame/types/hexvector"

//
// This class represents a set of hexes on a map.
//


type Region struct {
	Origin hexvector.Vector `bson:"origin"`
	Size hexvector.Vector `bson:"size"`
}

func (r *Region) Contains(v hexvector.Vector) (bool) {

	// Test if it's off the left or right side
	if v.Hx < r.Origin.Hx || v.Hx >= r.Origin.Hx + r.Size.Hx {
		return false
	}
	
	bias := ybias(v.Hx)
	if v.Hy < r.Origin.Hy + bias || v.Hy >= r.Origin.Hy + r.Size.Hy + bias {
		return false
	}
	
	return true
}

func (r *Region) All() (all []hexvector.Vector) {

	for hx := r.Origin.Hx ; hx < r.Size.Hx ; hx++ {
		for by := r.Origin.Hy ; by < r.Size.Hy ; by++ {
			hy := by + ybias(hx)
			all = append(all, hexvector.Vector{Hx: hx, Hy: hy})
		}
	}

	return
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

