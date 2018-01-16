package server

import (
	"github.com/markllama/hexgame/pkg/db"
)

func Main() {

	// Configure

	// connect to database
	database := db.Connect()

	CreateWebServer(database)
}
