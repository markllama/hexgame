package db

import (
	//"fmt"
	//"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Map struct {
	Name string `json:"name"`
	Game string `json:"game"`
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
