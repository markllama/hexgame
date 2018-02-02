package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"net/http"
	"net/url"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/markllama/hexgame/types/db"
)

//type MatchRef struct {
//	
//}

// MatchHandleFunc processes and responds to HTTP queries
// GET/name - return one
// GET/     - return references
// POST     - create a new one
func MatchHandleFunc(s *mgo.Session) (func(w http.ResponseWriter, r *http.Request)) {

	f := func(w http.ResponseWriter, r *http.Request) {
		
		sc := s.Copy()
		defer sc.Close()

		switch r.Method {
		// return an existing match or list of matches
		case http.MethodGet:
			_, id := path.Split(r.URL.Path)
			if (id == "") {
				GetMatchList(s, w, r)
			} else {
				GetMatch(s, id, w, r)
			}
		// create a new match
		case http.MethodPost:
			CreateMatch(s, w, r)

		// update/replace an existing match
		//case http.MethodPut:
		//	UpdateMatch(s, w, r)


		// delete a match
		//case http.MethodDelete:
		//	DeleteMatch(s, w, r)
		}
	}

	return f
}

type MatchRef struct {
	Id bson.ObjectId `json:"id"`
	GameId bson.ObjectId `json:"game_id"`
	URL string `json:"url"`
}

func GetMatchList(s *mgo.Session, w http.ResponseWriter, r *http.Request) {

	var m []db.Match

	c := s.DB("hexgame").C("matches")
	c.Find(nil).All(&m)

	matchrefs := make([]MatchRef, len(m))

	murl := url.URL{Scheme: "http", Host: r.Host}

	for index, match := range m {
		murl.Path = path.Join(r.URL.Path, match.Id.Hex())
		matchrefs[index].Id = match.Id
		matchrefs[index].GameId = match.GameId
		matchrefs[index].URL = murl.String()
	}
	
	jmatch, _ := json.Marshal(matchrefs)

	w.Write([]byte(jmatch))
}

func GetMatch(s *mgo.Session, id string, w http.ResponseWriter, r *http.Request) {
	var m db.Match

	c := s.DB("hexgame").C("matches")
	//c.Find(nil).All(&m)
	
	q := c.Find(bson.M{"_id": bson.ObjectIdHex(id)})
	// check for errors
	err := q.One(&m)
	if (err != nil) {
		http.Error(w, fmt.Sprintf("match %s not found", id), http.StatusNotFound)
		return
	}

	
	murl := url.URL{Scheme: "http", Host: r.Host, Path: r.URL.Path}
	m.URL = murl.String()
	p, _ := json.Marshal(m)
	w.Write(p)
}

func CreateMatch(s *mgo.Session, w http.ResponseWriter, r *http.Request) {

	// 
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var m db.Match

	// marshal the POST data into a Match
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	
	// Required: owner ID
	// Required: game ID
}

//func DeleteMatch(s *mgo.Session, w, http.ResponseWriter, r *http.Response) {
//
//}
