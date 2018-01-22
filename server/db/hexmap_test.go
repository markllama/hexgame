package db

import (
	"testing"
	gamedb "github.com/markllama/hexgame/pkg/db"
	"github.com/markllama/hexgame/types/hexmap"
)


func TestMapGet(t *testing.T) {
	db := gamedb.Connect()
	m := Map{db.C("maps"), hexmap.Map{Name: "clear"}, false}
	m.Get()
	m.Json()
}
