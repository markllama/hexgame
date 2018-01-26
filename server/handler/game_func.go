package handler

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

func GameServerFunc(s *mgo.Session) (func(http.ResponseWriter, *http.Request)) {


	f := func (w http.ResponseWriter, r *http.Request) {
		var p string

		session := s.Copy()
		defer session.Close()
	
		col := session.DB("hexgame").C("game")

		w.Header().Add("Content-Type", "application/json")

		_, name := path.Split(r.URL.Path)
		
		if (name != "") {
			g := db.Game{col, hexgame.Game{ Name : name }, false}
			err := g.Get()
			if (err != nil) {
				http.Error(w, fmt.Sprintf("game %s not found: %s", name, err), 404)
				return
			}

			gurl := url.URL{Scheme: "http", Host: r.Host, Path: r.URL.Path}
			g.URL = gurl.String()
			p = g.Json()
			//p = fmt.Sprintf("root = %s, name = %s", root, name)
			//p = fmt.Sprintf("root = %s, no name", root)
			w.Write([]byte(p))
		} else {

			games, _ := db.AllGames(col)
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

	return f
}
