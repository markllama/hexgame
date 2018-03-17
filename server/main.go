package server

import (
	"time"
//	"os"
//	"fmt"
	"strconv"
//	"html"
//	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	//"github.com/markllama/hexgame/server/handler"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

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

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/test", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"This is a test"})
		}
}


//func Main(opts *HexGameConfig) {

	// connect to database
	//session := Connect(&opts.MongoDBConfig)
//	Connect(&opts.MongoDBConfig)

	
	// http.Handle("/html/",
	// 	http.StripPrefix("/html/",
	// 		http.FileServer(http.Dir(opts.ContentRoot + "/html"))))
	
	// http.Handle("/js/",
	// 	http.StripPrefix("/js/",
	// 		http.FileServer(http.Dir(opts.ContentRoot + "/js"))))
	
	// http.HandleFunc("/api/game/", GameHandleFunc(session))
	// http.HandleFunc("/api/map/", MapHandleFunc(session))
	// http.HandleFunc("/api/match/", MatchHandleFunc(session))
	
	// http.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Good Bye: %s", html.EscapeString(r.URL.Path))
	// 	os.Exit(0)
	// })
	
	// log.Fatal(http.ListenAndServe(":8080", nil))
//}

