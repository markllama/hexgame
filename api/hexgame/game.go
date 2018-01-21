package hexgame

import (
	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	ID bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Title string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Copyright string `json:"copyright" bson:"copyright"`
	Description string `json:"description" bson:"description"`
	URL string `json:"url,omitempty" bson:"-"`
}
