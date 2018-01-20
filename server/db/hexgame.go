package db

import (
	"github.com/markllama/hexgame/api/hexgame"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

type Game struct {
	mcol *mgo.Collection `json:"-" bson:"-"`
	hexgame.Game
}

func (g *Game) Get() {
	q := g.mcol.Find(bson.M{"name": g.Name})	
	q.One(&g.Game)
}

func (g Game) Print() {
	fmt.Println(g)
}


