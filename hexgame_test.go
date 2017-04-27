package hexgame

import (
	"testing"
	"gopkg.in/mgo.v2"
	"net/http"
)

func TestHexGameStruct(t *testing.T) {
	session, err := mgo.Dial("mongodb://hexgame:ragnar@127.0.0.1/hexgame")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	gdb := session.DB("hexgame")

	gameListHandler := NewGameListHandler(gdb)
	gameHandler := NewGameHandler(gdb)

	http.HandleFunc("/games", gameListHandler)
	http.HandleFunc("/games/", gameHandler)
	http.ListenAndServe(":8000", nil)
}
