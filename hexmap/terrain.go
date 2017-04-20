package hexmap

type Terrain struct {
	Type string
	Name string
	Locations []HexVector `json:"locations" bson:"locations"`
}
