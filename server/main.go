package server

import (
	"time"
	"os"
	"fmt"
	"html"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	//"github.com/markllama/hexgame/server/handler"
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

	return session
}

func Main(code_root *string) {

	// connect to database
	session := Connect()

	http.Handle("/html/",
		http.StripPrefix("/html/",
			http.FileServer(http.Dir(*code_root + "/html"))))
	
	http.Handle("/js/",
		http.StripPrefix("/js/",
			http.FileServer(http.Dir(*code_root + "/js"))))
	
	http.HandleFunc("/api/game/", GameHandleFunc(session))
	http.HandleFunc("/api/map/", MapHandleFunc(session))
	http.HandleFunc("/api/match/", MatchHandleFunc(session))
	
	http.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Good Bye: %s", html.EscapeString(r.URL.Path))
		os.Exit(0)
	})
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

