package db

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Match struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	OwnerId bson.ObjectId `json:"owner_id" bson:"owner_id"`
	GameId bson.ObjectId `json:"game_id" bson:"game_id"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	StartTime time.Time `json:"start_time" bson:"start_time"`
	Players []bson.ObjectId `json:"players"`
	MapId bson.ObjectId `json:"map_id"`
	URL string `bson:"-" json:"url"`
}
