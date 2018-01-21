package handler

// when a caller requests /game/ or /game/name
// return the list of games or a specific named game

import (
	"net/http"
	"net/url"
	"gopkg.in/mgo.v2"
	"github.com/markllama/hexgame/types/hexgame"
	"github.com/markllama/hexgame/server/db"
	"fmt"
	"path"
	"encoding/json"
)

func GameServer(db *mgo.Database) (http.Handler) {
	return Game{mcol: db.C("games")}
}

type Game struct {
  mcol *mgo.Collection
}

type GameRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func (hg Game) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var p string
	
	w.Header().Add("Content-Type", "application/json")

	_, name := path.Split(r.URL.Path)

	if (name != "") {
		g := db.Game{hg.mcol, hexgame.Game{ Name : name }, false}
		err := g.Get()
		if (err != nil) {
			http.Error(w, fmt.Sprintf("game %s not found", name), 404)
			return
		}

		gurl := url.URL{Scheme: "http", Host: r.Host, Path: r.URL.Path}
		g.URL = gurl.String()
		p = g.Json()
		//p = fmt.Sprintf("root = %s, name = %s", root, name)
		//p = fmt.Sprintf("root = %s, no name", root)
		w.Write([]byte(p))
	} else {

		games, _ := db.AllGames(hg.mcol)
		gamerefs := make([]GameRef, len(games))
		
		gurl := url.URL{Scheme: "http", Host: r.Host}
		
		for index, game := range games {
			gurl.Path = path.Join(r.URL.Path, game.Name)
			gamerefs[index].Name = game.Name
			gamerefs[index].URL = gurl.String()
		}
		
		jgames, _ := json.Marshal(gamerefs)
		
		w.Write([]byte(jgames))
	}
}
