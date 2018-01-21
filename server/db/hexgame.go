package db

import (
	"github.com/markllama/hexgame/types/hexgame"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"fmt"
)

func AllGames(col *mgo.Collection) (games []Game, err error) {

	var hg []hexgame.Game
	q := col.Find(nil)


	n, err := q.Count()
	games = make([]Game, n)
	
	err = q.All(&hg)

	for index, game := range hg {
		games[index] = Game{Col: col, Game: game}
	}
	
	return
}

type Game struct {
	Col *mgo.Collection `json:"-" bson:"-"`
	hexgame.Game
	Clean bool `json:"-" bson:"-"`
}

func (g *Game) Get() (error) {
	q := g.Col.Find(bson.M{"name": g.Name})
	// check for errors
	err := q.One(&g.Game)
	if (err != nil) { g.Clean = true }
	return err
}

func (g Game) Print() {
	fmt.Println(g)
}

func (g Game) Json() (string) {
	jba, _ := json.Marshal(g)
	return string(jba)
}
