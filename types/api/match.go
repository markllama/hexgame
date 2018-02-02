package api

import (
	"gopkg.in/mgo.v2/bson"
)

type MatchRef struct {
	Id bson.ObjectId `json:"id"`
	GameId bson.ObjectId `json:"game_id"`
	URL string `json:"url"`
}
