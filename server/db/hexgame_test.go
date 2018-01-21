package db

import (
	"testing"
	gamedb "github.com/markllama/hexgame/pkg/db"
	"github.com/markllama/hexgame/api/hexgame"

)


func TestGameGet(t *testing.T) {
	db := gamedb.Connect()
	g := Game{db.C("games"), hexgame.Game{Name: "clear"}}
	g.Get()
	g.Json()
}
