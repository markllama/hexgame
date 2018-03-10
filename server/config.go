package server

import (
	"log"
	"fmt"
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


func LoadConfig(filename string) (*HexGameConfig) {
	
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


