package db

import (
	"fmt"
	//"encode/json"
	"gpkg.in/mgo.v2"
	"gpkg.in/mgo.v2/bson"
)


func SampleMaps(db) () {

	maps = db.C("maps")

	map1 = maps.Find(nil).One()

	fmt.PrintLn(map1)
}
