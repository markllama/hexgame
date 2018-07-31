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
	ContentRoot string `json:"content-root"`
	Port int `json:"port"`
	MongoDBConfig `json:"db-config,inline"`
}

type HexGameOptions struct {
	ConfigFile string
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

var defaults = HexGameOptions{
	ConfigFile: "/var/lib/hexgame/config.json",
	HexGameConfig: HexGameConfig{
		ContentRoot: "static",
		Port: 3000,
		MongoDBConfig: dbDefaults,
	},
}

// merge two configurations.  The first one takes precence. Only
// copy a value from c2 to c1 if the detination is nil.
func MergeConfigs(c0 *HexGameConfig, c1 *HexGameConfig) *HexGameConfig {

	if c0.ContentRoot == "" && c1.ContentRoot != "" {
		c0.ContentRoot = c1.ContentRoot
	}

	if c0.Port == 0 && c1.Port != 0 {
		c0.Port = c1.Port
	}

	if c0.DbServer == "" && c1.DbServer != "" {
		c0.DbServer = c1.DbServer
	}

	if c0.DbPort == 0 && c1.DbPort != 0 {
		c0.DbPort = c1.DbPort
	}

	if c0.DbName == "" && c1.DbName != "" {
		c0.DbName = c1.DbName
	}

	if c0.DbUser == "" && c1.DbUser != "" {
		c0.DbUser = c1.DbUser
	}

	if c0.DbPassword == "" && c1.DbPassword != "" {
		c0.DbPassword = c1.DbPassword
	}

	return c0
}

func GetConfig() *HexGameConfig {

	// retrieve config information from all three sources:
	cli := processFlags()
	env := processEnv()

	var config_file string
	
	if cli.ConfigFile != "" {
		config_file = cli.ConfigFile
	} else if env.ConfigFile != "" {
		config_file = env.ConfigFile
	} else {
		config_file = defaults.ConfigFile
	}
	
	fconf := loadConfig(config_file)

	// merge the inputs:
	// Precedence: cli > env > config > default

	var config *HexGameConfig
	config = MergeConfigs(&cli.HexGameConfig, &env.HexGameConfig)
	config = MergeConfigs(config, fconf)
	config = MergeConfigs(config, &defaults.HexGameConfig)
	
	return config
}


func processEnv() *HexGameOptions {

	var env HexGameOptions

	env.ConfigFile = os.Getenv("HEXGAME_CONFIG_FILE")
	env.ContentRoot = os.Getenv("HEXGAME_CONTENT_ROOT")
	port_string := os.Getenv("HEXGAME_PORT")
	if len(port_string) > 0 {
		var err error
		env.Port, err = strconv.Atoi(port_string)
		if err != nil {
			fmt.Printf("error parsing server port string '%s': %s\n", port_string, err)
		}
	}

	env.DbServer = os.Getenv("HEXGAME_DBSERVER")
	dbport_string := os.Getenv("HEXGAME_DBPORT")
	if len(dbport_string) > 0 {
		var err error
		env.DbPort, err = strconv.Atoi(dbport_string)
		if err != nil {
			fmt.Printf("error parsing DB port string '%s': %s\n", dbport_string, err)
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

	flag.StringVar(&opts.ConfigFile, "config-file", "",
		"The location of the hexgame server configuration file")

	flag.IntVar(&opts.Port, "port", 0,
		"The TCP port of the hexgame database")

	flag.StringVar(&opts.DbServer, "dbserver", "",
		"The hostname or IP address of a mongodb server holding the game database")

	flag.IntVar(&opts.DbPort, "dbport", 0,
		"The TCP port of the hexgame database")

	flag.StringVar(&opts.DbName, "dbname", "",
		"The name of the hexgame database")

	flag.StringVar(&opts.DbUser, "dbuser", "",
		"The access user for the hexgame database")

	flag.StringVar(&opts.DbPassword, "dbpass", "",
		"The access password for the hexgame database")

	flag.StringVar(&opts.ContentRoot, "content-root", "",
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
