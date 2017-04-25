package main

import (
	"log"
//	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"
)

type Config struct {
	DbName string `json:"db_name"`
	DbPassword string `json:"db_password"`
	DbServer string `json:"db_server"`
	DbUsername string `json:"db_username"`
}

func main() {

	// load configuration
	config_string, err := ioutil.ReadFile("hexmaprc.json")
	if err != nil {
		panic(err)
	}

	var config Config

	err = json.Unmarshal(config_string, &config)
	if err != nil {
		panic(err)
	}

	dbUrl := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		config.DbUsername,
		config.DbPassword,
		config.DbServer,
		config.DbName,
	)
	
	// initialize DB connection
	dbSession, err := mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	}
	defer dbSession.Close()
	dbSession.SetMode(mgo.Monotonic, true)

	http.HandleFunc("/games", GameHandler)
	
	// run web service
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world\n")
}
