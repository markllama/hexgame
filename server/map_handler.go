package server

import (
	"fmt"
	"path"
	"net/http"
	"net/url"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/markllama/hexgame/types/db"
	"github.com/markllama/hexgame/types/api"
)

func MapHandleFunc(s *mgo.Session) (func(w http.ResponseWriter, r *http.Request)) {

	f := func(w http.ResponseWriter, r *http.Request) {

		var md db.Map

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
			err := q.One(&md)
			if (err != nil) {
				http.Error(w, fmt.Sprintf("map %s not found", name), 404)
				return
			}

			var ma api.Map

			ma.Name = md.Name
			// ma.Game = db.Game.Get(md.GameId).Name
			ma.Copyright = md.Copyright
			ma.Shape = md.Shape
			ma.Copyright = md.Copyright
			ma.Size.Copy(md.Size)
			ma.Origin.Copy(md.Origin)
			
			murl := url.URL{Scheme: "http", Host: r.Host, Path: r.URL.Path}
			ma.URL = murl.String()
			p, _ := json.Marshal(ma)
			w.Write(p)

		//
		// Return references to all available maps
		//
		} else {

			var hm []db.Map
			c.Find(nil).All(&hm)

			maprefs := make([]api.MapRef, len(hm))

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
