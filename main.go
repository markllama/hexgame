// Copyright 2018 Mark Lamourine <markllama@gmail.com>
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
	"flag"
//	"os"
	server "github.com/markllama/hexgame/server"
)

func processFlags() *server.Options {

	var opts server.Options

	flag.StringVar(&opts.ConfigFile, "config-file", "hexgame-config.json",
		"The location of the hexgame server configuration file")
	
	flag.StringVar(&opts.DbServer, "dbserver", "localhost",
		"The hostname or IP address of a mongodb server holding the game database")

	flag.IntVar(&opts.DbPort, "dbport", 27017,
		"The TCP port of the hexgame database")

	flag.StringVar(&opts.DbName, "dbname", "hexgame",
		"The name of the hexgame database")

	flag.StringVar(&opts.DbUser, "dbuser", "hexgame",
		"The access user for the hexgame database")

	flag.StringVar(&opts.DbPassword, "dbpass", "hexgame",
		"The access password for the hexgame database")

	flag.StringVar(&opts.ContentRoot, "content-root", "./static",
		"The location of the static content for the game server")

	flag.BoolVar(&opts.Debug, "debug", false,
		"write debugging information")

	flag.BoolVar(&opts.Verbose, "verbose", false,
		"write verbose progress information")
	
	flag.Parse()
	return &opts
}

//func mergeEnvironment() {
//
//}

func main() {

	//cwd, _ := os.Getwd()
	//content_root := flag.String("content-root", cwd + "/static",
	//	"the location of the static content")

	opts := processFlags()
	
	server.Main(opts)
}
