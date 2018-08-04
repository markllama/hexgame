package api

import (
	"fmt"
	"path"
	"net/http"
	"net/url"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"

	"github.com/gorilla/mux"
	db "github.com/markllama/hexgame/types/db" 
	apitypes "github.com/markllama/hexgame/types/api" 
)


type MapsHandler struct {}

func (h MapsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this assumes that there is a mgo Session stored in the request context

	// get the hexgame database and games collection
	context := r.Context()
	session := context.Value("mongoSession").(*mgo.Session)
	maps := session.DB("hexgame").C("maps")

	var hm []db.Map
	
	maps.Find(nil).All(&hm)

	maprefs := make([]apitypes.MapRef, len(hm))

	w.Header().Add("Content-Type", "application/json")
	murl := url.URL{Scheme: "http", Host: r.Host}
		
	for index, hmap := range hm {
		murl.Path = path.Join(r.URL.Path, hmap.Name)
		maprefs[index].Name = hmap.Name
		maprefs[index].URL = murl.String()
	}
		
	jmaps, _ := json.Marshal(maprefs)
	//			jgames, _ := json.Marshal(hg)
	
	w.Write([]byte(jmaps))
}


type MapHandler struct {}

func (h MapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this assumes that there is a mgo Session stored in the request context

	// get the hexgame database and games collection
	context := r.Context()
	session := context.Value("mongoSession").(*mgo.Session)
	maps := session.DB("hexgame").C("maps")

	rvars := mux.Vars(r)
	mapName := rvars["name"]
	
	var hm db.Map
	
	q := maps.Find(bson.M{"name": mapName})
	err := q.One(&hm)

	if (err != nil) {
		http.Error(w, fmt.Sprintf("map %s not found", mapName), 404)
		return
	}

	// get the game ID and retrieve the game name
	gameId := hm.GameId

	games := session.DB("hexgame").C("games")
	var hg db.Game

	gq := games.Find(bson.M{"_id": gameId})
	gerr := gq.One(&hg)
	
	if (gerr != nil) {
		http.Error(w, fmt.Sprintf("game %s not found", gameId), 404)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	murl := url.URL{Scheme: "http", Host: r.Host}

	apiMap := new(apitypes.Map)
	
	murl.Path = path.Join(r.URL.Path, apiMap.Name)
	apiMap.Name = hm.Name
	apiMap.URL = murl.String()
	apiMap.Game = hg.Name
	apiMap.Copyright = hm.Copyright
	apiMap.Shape = hm.Shape
	apiMap.Size = hm.Size
	apiMap.Origin = hm.Origin
	apiMap.Terrains = hm.Terrains
		
	jmap, _ := json.Marshal(apiMap)
	
	w.Write([]byte(jmap))
}
