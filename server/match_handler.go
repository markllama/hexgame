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
	"github.com/markllama/hexgame/types/api"
)

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

func GetMatchList(s *mgo.Session, w http.ResponseWriter, r *http.Request) {

	var m []db.Match

	c := s.DB("hexgame").C("matches")
	c.Find(nil).All(&m)

	matchrefs := make([]api.MatchRef, len(m))

	murl := url.URL{Scheme: "http", Host: r.Host}

	for index, match := range m {
		murl.Path = path.Join(r.URL.Path, match.Id.Hex())
		matchrefs[index].Id = match.Id

		// get the game with the specified ID and retrieve the name
		matchrefs[index].Game = match.GameId.Hex()
		matchrefs[index].URL = murl.String()
	}
	
	jmatch, _ := json.Marshal(matchrefs)

	w.Write([]byte(jmatch))
}

func GetMatch(s *mgo.Session, id string, w http.ResponseWriter, r *http.Request) {
	var m db.Match

	c_matches := s.DB("hexgame").C("matches")
	//c.Find(nil).All(&m)
	
	q := c_matches.Find(bson.M{"_id": bson.ObjectIdHex(id)})
	// check for errors
	err := q.One(&m)
	if (err != nil) {
		http.Error(w, fmt.Sprintf("match %s not found", id), http.StatusNotFound)
		return
	}

	c_games := s.DB("hexgame").C("games")
	var g db.Game
	err = g.Get(c_games, bson.M{"_id": m.GameId})
	if (err != nil) {
		http.Error(w, fmt.Sprintf("game not found"), http.StatusInternalServerError)
		return
	}

	c_users := s.DB("hexgame").C("users")
	var u db.User
	err = u.Get(c_users, bson.M{"_id": m.OwnerId})
	if (err != nil) {
		http.Error(w, fmt.Sprintf("user not found"), http.StatusInternalServerError)
		return
	}

	var data api.Match
	data.Id = m.Id
	data.Game = g.Name
	data.Owner = u.Username
	data.CreateTime = m.CreateTime
	data.StartTime = m.StartTime

	murl := url.URL{Scheme: "http", Host: r.Host, Path: r.URL.Path}
	data.URL = murl.String()
	p, _ := json.Marshal(data)
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
