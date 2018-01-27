package server

import (
	"os"
	"fmt"
	"html"
	"log"	
	"net/http"
	"gopkg.in/mgo.v2"

	"github.com/markllama/hexgame/server/handler"
)

func CreateWebServer(s *mgo.Session) {

	db := s.DB("hexgame")

	http.HandleFunc("/game/", handler.GameServerFunc(s))
	//http.Handle("/game/", handler.GameServer(db))
	http.Handle("/map/", handler.MapServer(db))
	
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
