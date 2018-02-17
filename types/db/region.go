package db


type Region struct {
	
	Shape string `bson:"shape"`
	Origin Vector `bson:"origin"`
	Size Vector `bson:"size"`
}
