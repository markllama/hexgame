package server

import (
	"flag"
	"log"
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type MongoDBConfig struct {
	DbServer string `json:"server"`
	DbPort int `json:"port"`
	DbName string `json:"database"`
	DbUser string `json:"username"`
	DbPassword string `json:"password"`
}

type  HexGameConfig struct {
	ConfigFile string `json:"-"`
	ContentRoot string `json:"content-root"`
	MongoDBConfig `json:"db-config,inline"`
}

type HexGameOptions struct {
	HexGameConfig
	Verbose bool
	Debug bool
}

var dbDefaults = MongoDBConfig{
	DbServer: "localhost",
	DbPort: 27017,
	DbName: "hexgame",
	DbUser: "hexgame",
}

var defaults = HexGameConfig{
	ConfigFile: "hexgame.json",
	ContentRoot: "static",
	MongoDBConfig: dbDefaults,
}

// merge two configurations.  The first one takes precence. Only
// copy a value from c2 to c1 if the detination is nil.
func MergeConfigs(c1 *HexGameConfig, c2 *HexGameConfig) *HexGameConfig {

	return c1
}

func GetConfig() *HexGameConfig {

	// retrieve config information from all three sources:
//	env := processEnv()
	cli := processFlags()
	config := loadConfig(cli.ConfigFile)

	// merge the inputs:
	// Precedence: cli > env > config > default

	
	
	
	return config
}


func processEnv() *HexGameOptions {

	var env HexGameOptions

	env.ConfigFile = os.Getenv("HEXGAME_CONFIG_FILE")
	env.ContentRoot = os.Getenv("HEXGAME_CONTENT_ROOT")

	env.DbServer = os.Getenv("HEXGAME_DBSERVER")
	port_string := os.Getenv("HEXGAME_DBPORT")
	if len(port_string) > 0 {
		var err error
		env.DbPort, err = strconv.Atoi(port_string)
		if err != nil {
			fmt.Printf("error parsing DB port string '%s': %s\n", port_string, err)
		}
	}
	env.DbName = os.Getenv("HEXGAME_DBNAME")
	env.DbUser = os.Getenv("HEXGAME_DBUSER")
	env.DbPassword = os.Getenv("HEXGAME_DBPASS")

	verbose := os.Getenv("HEXGAME_VERBOSE")
	if len(verbose) > 0 {
		env.Verbose, _ = strconv.ParseBool(verbose)
	}

	debug := os.Getenv("HEXGAME_DEBUG")
	if len(debug) > 0 {
		env.Debug, _ = strconv.ParseBool(debug)
	}

	return &env
}

func processFlags() *HexGameOptions {

	var opts HexGameOptions

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

func loadConfig(filename string) (*HexGameConfig) {
	
	var config HexGameConfig
	
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(
			fmt.Sprintf(
				"unable to open test config file -  %s: %s", filename, err))
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(
			fmt.Sprintf(
				"unable to marshal config -  %s: %s", filename, err))
	}

	return &config
}
