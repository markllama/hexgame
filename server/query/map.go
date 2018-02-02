package query

import (
	"github.com/markllama/hexgame/types/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	//"fmt"
)

type Map struct {
	Col *mgo.Collection `json:"-" bson:"-"`
	db.Map
	Clean bool `json:"-" bson:"-"`
}

func AllMaps(col *mgo.Collection) (maps []Map, err error) {
	var hm []db.Map
	q := col.Find(nil)

	n, err := q.Count()
	maps = make([]Map, n)

	err = q.All(&hm)

	for index, hmap := range hm {
		maps[index] = Map{Col: col, Map: hmap}
	}

	return
}

func (m *Map) Get() (error) {
	q := m.Col.Find(bson.M{"name": m.Name})
	// check for errors
	err := q.One(&m.Map)
	if (err != nil) { m.Clean = true }
	return err
}

func (m Map) Json() (string) {
	jba, _ := json.Marshal(m)
	return string(jba)
}
