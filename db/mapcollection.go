package db

import (
	"github.com/markllama/hexgame/types/hexmap"
)

type hexmapCollection interface {
	addHexmap(hm hexmap.Map) (err error)
	getHexmaps() []hexmap.Map
	getHexmap(id string) (hm hexmap.Map, err error)
}
