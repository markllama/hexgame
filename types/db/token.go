package db

import (
	"github.com/markllama/hexgame/types"
)


type Token struct {
	Name string `json:"name"`
	Location types.Vector `json:"location"`
}
