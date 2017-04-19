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

	c := session.DB("hexgame").C("maps")
	err = c.Insert(hexmap.HexMap{"Ogre", *hexmap.NewHexVector(22,14), hexmap.ORIGIN}, hexmap.HexMap{"WarpWar", *hexmap.NewHexVector(22, 15), hexmap.ORIGIN})
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
