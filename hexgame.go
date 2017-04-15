package main


import (
	"fmt";
	"log";
	"gopkg.in/mgo.v2";
	"gopkg.in/mgo.v2/bson";
	//"net/http";
)

type Hexmap struct {
	Name string
	Size string
}

func main() {
	fmt.Println("Hello there")

	session, err := mgo.Dial("mongodb://hexgame:ragnar@172.17.0.2/hexgame")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("hexgame").C("maps")
	err = c.Insert(&Hexmap{"Ogre", "14x22"}, &Hexmap{"WarpWar", "14x24"})
	if err != nil {
		log.Fatal(err)
	}

	result := Hexmap{}
	err = c.Find(bson.M{"name": "Ogre"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Hexmap:", result.Name)
}
