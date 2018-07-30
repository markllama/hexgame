/// Copyright 2018 Mark Lamourine <markllama@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"net/http"
	//"log"
	"time"
	config "github.com/markllama/hexgame/config"
	db "github.com/markllama/hexgame/db"
)

type HttpHandlerDecorator func(http.Handler) http.Handler

// the configuration is integrated from defaults, environment, config file
// and CLI arguments in that order of precedence.
func main() {

	opts := config.GetConfig()
	// read configuration

	// connect to database
	session := db.Connect(&opts.MongoDBConfig)

	fmt.Println(session)
	// create api server

	apiHandler := http.FileServer(http.Dir("./static"))
	apiServer := &http.Server{
		Addr:           ":8080",
		Handler:        apiHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	dbDecorator := db.CopyMongoSession(session)
	
	appHandler := dbDecorator(http.FileServer(http.Dir("./static")))
	appServer := &http.Server{
		Addr:           ":8999",
		Handler:        appHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go apiServer.ListenAndServe()

	// create user interface server
	appServer.ListenAndServe()

	fmt.Printf("hello world %s", opts)
}
