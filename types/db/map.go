package db

import (
	//"fmt"
	//"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Map struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string `json:"name"`
	GameId bson.ObjectId `json:"game_id" bson:"game_id"`
	Copyright string `json:"copyright"`
	Shape string `json:"shape"`
	Size Vector `json:"size"`
	Origin Vector `json:"origin"`
	Terrains []Terrain `json:"terrains"`
	Tokens []Token `json:"tokens,omitempty"`
}

func (m *Map) Get(c mgo.Collection, selector bson.M) (error) {
	return c.Find(selector).One(&m)
}

