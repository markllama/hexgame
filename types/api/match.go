package api

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Match struct {
	Id bson.ObjectId `json:"id"`	
	CreateTime time.Time `json:"create_time"`
	StartTime time.Time `json:"start_time"`
	Game string `json:"game"`
	Owner string `json:"owner"`
	Players []string `json:"players"`
	Map string `json:"map"`
	URL string `json:"url"`
}

type MatchRef struct {
	Id bson.ObjectId `json:"id"`
	Game string `json:"game"`
	URL string `json:"url"`
}
