package server

import (
	"time"
	"os"
	"fmt"
	"html"
	"log"
	"path"
	"net/http"
	"net/url"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/markllama/hexgame/types/hexgame"
	"github.com/markllama/hexgame/server/handler"
)


//const MongoDb details
const (
	hosts      = "127.0.0.1:27017"
	database   = "hexgame"
	username   = "hexgame"
	password   = "ragnar"
)

func Connect() (*mgo.Session) {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	//db := session.DB(database)
	//return db

	return session
}

func Main() {

	// connect to database
	session := Connect()
	database := session.DB("hexgame")
	
	//http.Handle("/game/", handler.GameServer(database))
	http.HandleFunc("/game/", GameHandleFunc(session))
	http.Handle("/map/", handler.MapServer(database))
	
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"foo\": \"bar\"}\n"))
	})

	http.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Good Bye", html.EscapeString(r.URL.Path))
		os.Exit(0)
	})
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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
			//p = fmt.Sprintf("root = %s, name = %s", root, name)
			//p = fmt.Sprintf("root = %s, no name", root)
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
