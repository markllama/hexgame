package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/markllama/hexgame/types"
)

type Map struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	GameId bson.ObjectId `bson:"game_id,omitempty"`
	types.Map `bson:",inline"`
}

func (m *Map) Get(c mgo.Collection, selector bson.M) (error) {
	return c.Find(selector).One(&m)
}

