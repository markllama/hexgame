package server

import (
	"fmt"
	"path"
	"net/http"
	"net/url"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/markllama/hexgame/types/hexgame"
)


type GameRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func GameHandleFunc(s *mgo.Session) (func(http.ResponseWriter, *http.Request)) {
	
	f := func(w http.ResponseWriter, r *http.Request) {
		var g hexgame.Game

		sc := s.Copy()
		defer sc.Close()

		c := sc.DB("hexgame").C("games")
		
		w.Header().Add("Content-Type", "application/json")

		_, name := path.Split(r.URL.Path)

		if (name != "") {

			q := c.Find(bson.M{"name": name})
			// check for errors
			// err := q.One(&g.Game)
			err := q.One(&g)
			if (err != nil) {
				http.Error(w, fmt.Sprintf("game %s not found", name), 404)
				return
			}

			gurl := url.URL{Scheme: "http", Host: r.Host, Path: r.URL.Path}
			g.URL = gurl.String()
			p, _ := json.Marshal(g)
			w.Write(p)
		} else {


			var hg []hexgame.Game
			q := c.Find(nil)

			n, _ := q.Count()
			games := make([]hexgame.Game, n)
			q.All(&hg)
	
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
