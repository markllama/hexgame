package server

import (
	"fmt"

	"github.com/markllama/hexgame/pkg/db"

)

func Main() {

	// Configure

	// connect to database
	database := db.Connect()

	// start web server
	fmt.Println(database)

	db.SampleMaps(database)
	
	
	fmt.Println("Hey there")


	//CreateWebServer()
}
