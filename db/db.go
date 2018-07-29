package db

import (
	"strconv"
	"time"
	
	"gopkg.in/mgo.v2"
	config "github.com/markllama/hexgame/config"

	"context"
	"net/http"
)

func Connect(opts *config.MongoDBConfig) (*mgo.Session) {

	host_port := opts.DbServer + ":" + strconv.Itoa(opts.DbPort)
	
	info := &mgo.DialInfo{
		
		Addrs:    []string{host_port},
		Timeout:  60 * time.Second,
		Database: opts.DbName,
		Username: opts.DbUser,
		Password: opts.DbPassword,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return session
}

type HttpHandlerDecorator func(http.Handler) http.Handler

//
func CopyMongoSession(db *mgo.Session) HttpHandlerDecorator {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// copy the database session
			dbsession := db.Copy()
			defer dbsession.Close() // clean up

			ctx := context.WithValue(r.Context(), "database", dbsession)
			// save it in the mux context
			new_r := r.WithContext(ctx)

			// pass execution to the original handler
			h.ServeHTTP(w, new_r)

		})
	}
}
