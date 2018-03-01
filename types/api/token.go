package api

import (
	"github.com/markllama/hexgame/types/hexvector"
)

type Token struct {
	Name string `json:"name"`
	Location hexvector.Vector `json:"location"`
}
