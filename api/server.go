package api

import (
	"time"
	"net/http"
	
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	db "github.com/markllama/hexgame/db"
)

func NewApiServer(dbSession *mgo.Session) (*http.Server) {

	dbDecorator := db.CopyMongoSession(dbSession)

	apiMux := mux.NewRouter()
	//apiMux := http.NewServeMux()

	apiMux.Handle("/games/", dbDecorator(new(GamesHandler)))
	apiMux.Handle("/maps/", dbDecorator(new(MapsHandler)))
	apiMux.Handle("/maps/{name}", dbDecorator(new(MapHandler)))

	apiServer := &http.Server{
		Addr:           ":8999",
		Handler:        apiMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return apiServer
}
