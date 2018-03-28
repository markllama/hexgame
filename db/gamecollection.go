package db

import (
	"github.com/markllama/hexgame/types"
)

type gameCollection interface {
	getGames() []types.Game
	getGame(id string) (game types.Game, err error)
	addGame(game types.Game) (err error)
}


