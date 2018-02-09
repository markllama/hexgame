package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/markllama/hexgame/types"
)

type Game struct {
	Id bson.ObjectId `json:"-" bson:"_id,omitempty"`
	types.Game `bson:",inline"`
}

func (g *Game) Get(c *mgo.Collection, selector bson.M) (error) {
	return c.Find(selector).One(&g)
}
