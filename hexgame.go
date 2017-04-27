package hexgame

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	//"net/url"
	//"io"
	"fmt"
	"strings"
	"encoding/json"
	"path"
)

type HexGame struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Title string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Copyright string `json:"copyright" bson:"copyright"`
	Description string `json:"description" bson:"description"`
}

// inject a DB session and get back an http handler
//
func NewGameListHandler(db *mgo.Database) (handler func(http.ResponseWriter, *http.Request)) {

	gameCollection := db.C("games")
	
	handler = func(w http.ResponseWriter, r *http.Request) {

		// create a filter if there are any query parameters
		filter := bson.M{}
		if val, present := r.URL.Query()["name"] ; present {
			fmt.Println("finding " + val[0])
			filter = bson.M{"name": val[0]}
		}
		
		var result []HexGame
		iter := gameCollection.Find(filter).Limit(100).Iter()
		err := iter.All(&result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			responses, err := json.Marshal(result)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")			
			fmt.Fprint(w, string(responses))
		}
	}

	return
}

func NewGameHandler(db *mgo.Database) (handler func(http.ResponseWriter, *http.Request)) {

	gameCollection := db.C("games")
	
	handler = func(w http.ResponseWriter, r *http.Request) {

		id_string := path.Base(r.URL.Path)
		if ! bson.IsObjectIdHex(path.Base(id_string)) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			id := bson.ObjectIdHex(id_string)
			var result HexGame
			query := gameCollection.FindId(id)
			err := query.One(&result)
			if err != nil {
				if strings.Compare(err.Error(), "not found") == 0 {
					w.WriteHeader(http.StatusNotFound)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				responses, err := json.Marshal(result)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, string(responses))
			}
		}
	}
	return
}


