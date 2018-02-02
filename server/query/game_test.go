package query

import (
	"testing"
	gamedb "github.com/markllama/hexgame/pkg/db"
	"github.com/markllama/hexgame/types/hexgame"

)


func TestGameGet(t *testing.T) {
	db := gamedb.Connect()
	g := Game{db.C("games"), hexgame.Game{Name: "clear"}, false}
	g.Get()
	g.Json()
}
