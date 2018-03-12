package server

import (
	"time"
	"os"
	"fmt"
	"strconv"
	"html"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	//"github.com/markllama/hexgame/server/handler"
)

type Options struct {

	ConfigFile string
	
	DbServer string
	DbPort int
	DbName string
	DbUser string
	DbPassword string

	ContentRoot string

	Debug bool
	Verbose bool
}

//const MongoDb details
var Defaults = Options{
	ConfigFile: "./hexgame_config.yaml",
	DbServer: "127.0.0.1",
	DbPort: 27017,
	DbName: "hexgame",
	DbUser: "hexgame",
	ContentRoot: "./static",

	Debug: false,
	Verbose: false,
}

func Environment() (opts *Options) {

	config_file := os.Getenv("HEXGAME_CONFIG_FILE")
	if len(config_file) > 0 {
		opts.ConfigFile = config_file
	}
	
	server := os.Getenv("HEXGAME_SERVER")
	if len(server) > 0 {
		opts.DbServer = server
	}

	port_str := os.Getenv("HEXGAME_PORT")
	if len(port_str) > 0 {
		port, err := strconv.Atoi(port_str)
		if err == nil {
			opts.DbPort = port
		}
	}

	user := os.Getenv("HEXGAME_USER")
	if len(user) > 0 {
		opts.DbUser = user
	}

	content_root := os.Getenv("HEXGAME_CONTENT_ROOT")
	if len(content_root) > 0 {
		opts.ContentRoot = content_root
	}

	return
}

func Connect(opts *MongoDBConfig) (*mgo.Session) {

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

func Main(opts *HexGameConfig) {

	// connect to database
	session := Connect(&opts.MongoDBConfig)

	http.Handle("/html/",
		http.StripPrefix("/html/",
			http.FileServer(http.Dir(opts.ContentRoot + "/html"))))
	
	http.Handle("/js/",
		http.StripPrefix("/js/",
			http.FileServer(http.Dir(opts.ContentRoot + "/js"))))
	
	http.HandleFunc("/api/game/", GameHandleFunc(session))
	http.HandleFunc("/api/map/", MapHandleFunc(session))
	http.HandleFunc("/api/match/", MatchHandleFunc(session))
	
	http.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Good Bye: %s", html.EscapeString(r.URL.Path))
		os.Exit(0)
	})
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

