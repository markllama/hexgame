package main

import (
	"flag"
	"log"
//	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/markllama/hexgame-server/hexgame"
)

type Config struct {
	DbName string `json:"db_name"`
	DbPassword string `json:"db_password"`
	DbServer string `json:"db_server"`
	DbUsername string `json:"db_username"`
}

func main() {

	var config_file string

	flag.StringVar(&config_file, "config-file", "hexgamerc.json",
		"Hexgame server configuration file")
	flag.Parse()
	
	// load configuration
	config_string, err := ioutil.ReadFile(config_file)
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
	gdb := dbSession.DB("hexgame")

	gameListHandler := hexgame.NewGameListHandler(gdb)
	gameHandler := hexgame.NewGameHandler(gdb)

	http.HandleFunc("/games", gameListHandler)
	http.HandleFunc("/games/", gameHandler)
	// http.HandleFunc("/games/{name}/maps", mapHandler)

	// run web service
	log.Fatal(http.ListenAndServe(":8080", nil))
}

