package main


import (
	"fmt";
	"log";
	"gopkg.in/mgo.v2";
	"gopkg.in/mgo.v2/bson";
	//"net/http";

	"github.com/markllama/hexgame-go/hexmap";
)


func main() {
	fmt.Println("Hello there")

	session, err := mgo.Dial("mongodb://hexgame:ragnar@172.17.0.2/hexgame")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	ogremap := hexmap.HexMap{
		Name: "Ogre",
		Size: *hexmap.NewHexVector(14,22),
		Origin: hexmap.ORIGIN,
		Terrains: []hexmap.Terrain{
			{
				"clear", "clear", []hexmap.HexVector{
					*hexmap.NewHexVector(3, 4),
					*hexmap.NewHexVector(2, 3),
					*hexmap.NewHexVector(4, 20),
				},
			},
		},
	}
		

	warpwarmap := hexmap.HexMap{
		Name: "WarpWar",
		Size: *hexmap.NewHexVector(15,22),
		Origin: hexmap.ORIGIN,
		Terrains: []hexmap.Terrain{
			{
				"clear", "clear", []hexmap.HexVector{
					*hexmap.NewHexVector(3, 4),
					*hexmap.NewHexVector(2, 3),
					*hexmap.NewHexVector(4, 20),
				},
			},
		},
	}

	c := session.DB("hexgame").C("maps")
	err = c.Insert(ogremap, warpwarmap)
	if err != nil {
		log.Fatal(err)
	}

	result := hexmap.HexMap{}
	err = c.Find(bson.M{"name": "Ogre"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Hexmap:", result.Name)
}
