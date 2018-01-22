package db

import (
	"fmt"
	//"encode/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/markllama/hexgame/types/hexgame"
)

func SampleGame(db *mgo.Database) () {

	games := db.C("games")

	result := hexgame.Game{}
	err := games.Find(bson.M{"name": "clear"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
