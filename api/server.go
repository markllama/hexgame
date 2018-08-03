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

	var gh GameHandler
	
	apiHandler := dbDecorator(gh)
	apiMux.Handle("/games/", apiHandler)

	apiServer := &http.Server{
		Addr:           ":8999",
		Handler:        apiMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return apiServer
}
