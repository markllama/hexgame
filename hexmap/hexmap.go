package hexmap

import "gopkg.in/mgo.v2/bson";

//import "encoding/json"

type HexMap struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	GameID bson.ObjectId `json:"_gameid,omitempty" bson:"_gameid,omitempty"`
	Name string `json:"name"`
	Shape string `json:"shape"`
	Size HexVector `json:"size"`
	Origin HexVector `json:"origin"`
	Terrains []Terrain `json:"terrains"`
}

//
// A triangular lattice is biased wrt a rectangular coordinate system
// every second column, the row is positive biased by 1
//
func (hm HexMap) ybias(hx int) int {
	// normalize to the x origin
	hn := hx - hm.Origin.Hx()

	if hn >= 0 { return hn / 2 }
	return -(abs(hn) + 1) / 2
}

// does the map contain a specific location
func (hm HexMap) Contains(hv HexVector) bool {
	// if the x value is out of range then definitely no

	if hv.Hx() < hm.Origin.Hx() { return false }
	if hv.Hx() > hm.Origin.Hx() + hm.Size.Hx() - 1 { return false }

	ybias := hm.ybias(hv.Hx())

	if hv.Hy() < hm.Origin.Hy() + ybias { return false }
	if hv.Hy() > hm.Origin.Hy() + hm.Size.Hy() + ybias - 1 { return false }
	return true
}


