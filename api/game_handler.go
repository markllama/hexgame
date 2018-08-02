package api

import (
	"path"
	"net/http"
	"net/url"
	"gopkg.in/mgo.v2"
	"encoding/json"


	db "github.com/markllama/hexgame/types/db" 
	apitypes "github.com/markllama/hexgame/types/api" 
)


type GameHandler struct {}

func (h GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this assumes that there is a mgo Session stored in the request context

	// get the hexgame database and games collection
	context := r.Context()
	session := context.Value("mongoSession").(*mgo.Session)
	games := session.DB("hexgame").C("games")

	var hg []db.Game
	
	games.Find(nil).All(&hg)

	gamerefs := make([]apitypes.GameRef, len(hg))

	w.Header().Add("Content-Type", "application/json")
	gurl := url.URL{Scheme: "http", Host: r.Host}
		
	for index, game := range hg {
		gurl.Path = path.Join(r.URL.Path, game.Name)
		gamerefs[index].Name = game.Name
		gamerefs[index].URL = gurl.String()
	}
		
	jgames, _ := json.Marshal(gamerefs)
	//			jgames, _ := json.Marshal(hg)
	
	w.Write([]byte(jgames))
}
