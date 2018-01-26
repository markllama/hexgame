package db

import (
	"time"
	"gopkg.in/mgo.v2"
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

	session, err1 := mgo.DialWithInfo(info)
	if err1 != nil {
		panic(err1)
	}

	//db := session.DB(database)
	
	return session
}
