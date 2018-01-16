package main

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string `json:"name"`
	Title string `json:"title"`
	Author string `json:"author"`
	Copyright string `json:"copyright"`
	Description string `json:"description"`
}

//const MongoDb details
const (
	hosts      = "127.0.0.1:27017"
	database   = "hexgame"
	username   = "hexgame"
	password   = "ragnar"
)

func Connect() (*mgo.Database) {
	fmt.Println("In db connect")

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err1 := mgo.DialWithInfo(info)
	if err1 != nil {
		panic(err1)
	}

	db := session.DB(database)
	
	return db
}

func main() {
	var result Game


	db := Connect()

	fmt.Println(db)

	err := db.C("games").Find(bson.M{"name": "warpwar"}).One(&result)
	if err != nil { panic (err) }
	
	fmt.Println(result)
	
	fmt.Println("Hello There")
}
