package handler

import (
	"net/http"
	"net/url"
	"gopkg.in/mgo.v2"
	"github.com/markllama/hexgame/types/db"
	"github.com/markllama/hexgame/server/query"
	"fmt"
	"path"
	"encoding/json"
)

func MapServer(db *mgo.Database) (http.Handler) {
	return Map{mcol: db.C("maps")}
}

type Map struct {
	mcol *mgo.Collection
}

type MapRef struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func (hm Map) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var p string
	w.Header().Add("Content-Type", "application/json")

	_, name := path.Split(r.URL.Path)

	if (name != "") {
		m := query.Map{hm.mcol, db.Map{Name: name}, false}
		err := m.Get()
		if (err != nil) {
			http.Error(w, fmt.Sprintf("map %s not found", name), 404)
			return
		}

		murl := url.URL{Scheme: "http", Host: r.Host, Path: r.URL.Path}
		m.URL = murl.String()
		p = m.Json()
		w.Write([]byte(p))
	} else {

		maps, _ := query.AllMaps(hm.mcol)
		maprefs := make([]MapRef, len(maps))

		murl := url.URL{Scheme: "http", Host: r.Host}

		for index, hmap := range maps {
			murl.Path = path.Join(r.URL.Path, hmap.Name)
			maprefs[index].Name = hmap.Name
			maprefs[index].URL = murl.String()
		}

		jmaps, _ := json.Marshal(maprefs)

		w.Write([]byte(jmaps))
	}
}
