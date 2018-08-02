package api

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2"

	db "github.com/markllama/hexgame/db"
)

func NewApiServer(dbSession *mgo.Session) (*http.Server) {

	dbDecorator := db.CopyMongoSession(dbSession)

	apiMux := http.NewServeMux()

	var gh GameHandler
	
	apiHandler := dbDecorator(gh)
	//apiHandler := http.FileServer(http.Dir("./static"))
	apiMux.Handle("/", apiHandler)

	apiServer := &http.Server{
		Addr:           ":8999",
		Handler:        apiMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return apiServer
}
