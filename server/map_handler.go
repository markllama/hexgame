package server

import (
	"fmt"
	"path"
	"net/http"
	"net/url"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/markllama/hexgame/types/hexmap"
)

type MapRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func MapHandleFunc(s *mgo.Session) (func(w http.ResponseWriter, r *http.Request)) {

	f := func(w http.ResponseWriter, r *http.Request) {

		var m hexmap.Map

		sc := s.Copy()
		defer sc.Close()

		c := sc.DB("hexgame").C("maps")

		w.Header().Add("Content-Type", "application/json")

		_, name := path.Split(r.URL.Path)

		//
		// Return a specific map if named
		//
		if (name != "") {
			q := c.Find(bson.M{"name": name})
			// check for errors
			err := q.One(&m)
			if (err != nil) {
				http.Error(w, fmt.Sprintf("map %s not found", name), 404)
				return
			}

			murl := url.URL{Scheme: "http", Host: r.Host, Path: r.URL.Path}
			m.URL = murl.String()
			p, _ := json.Marshal(m)
			w.Write(p)

		//
		// Return references to all available maps
		//
		} else {

			var hm []hexmap.Map
			c.Find(nil).All(&hm)

			maprefs := make([]MapRef, len(hm))

			murl := url.URL{Scheme: "http", Host: r.Host}

			for index, hmap := range hm {
				murl.Path = path.Join(r.URL.Path, hmap.Name)
				maprefs[index].Name = hmap.Name
				maprefs[index].URL = murl.String()
			}

			jmaps, _ := json.Marshal(maprefs)

			w.Write([]byte(jmaps))
		}

	}
	
	return f
}
