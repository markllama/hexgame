package db

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Match struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	OwnerId bson.ObjectId `json:"owner_id" bson:"owner_id,omitempty"`
	GameId bson.ObjectId `json:"game_id" bson:"game_id,omitempty"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	Players []bson.ObjectId `json:"players"`
	MapId bson.ObjectId `json:"map_id bson:"map_id,omitempty"`
}

func (m *Match) Get(c *mgo.Collection, selector bson.M) (error) {
	return c.Find(selector).One(&m)
}

func (m Match) Put(c *mgo.Collection) (error) {
	return c.Insert(&m)
}
