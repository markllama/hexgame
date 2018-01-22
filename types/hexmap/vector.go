package hexmap

import (
	"gopkg.in/mgo.v2"
)

type Vector struct {
	Hx int `json:"hx" bson:"hx"`
	Hy int `json:"hx" bson:"hx"`
}
