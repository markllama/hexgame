package server

import (
	"testing"

	"fmt"
	"os"
)

func TestLoad(t *testing.T) {

	test_config_file := "../test/hexgame_config.json"

	config := loadConfig(test_config_file)

	expected := MongoDBConfig{
		DbServer: "localhost",
		DbPort: 27017,
		DbName: "hexgame",
		DbUser: "hexgame",
		DbPassword: "ragnar",
	}
	
	if config.DbServer != expected.DbServer {
		t.Error(
			fmt.Sprintf("DbServer does not match: expected = %s, actual %s",
				expected.DbServer,
				config.DbServer))
	}

	if config.DbPort != expected.DbPort {
		t.Error(
			fmt.Sprintf("DbPort does not match: expected %d, actual %d",
				expected.DbPort,
				config.DbPort))
	}

	if config.DbName != expected.DbName {
		t.Error(
			fmt.Sprintf("DbName does not match: expected = %s, actual %s",
				expected.DbName,
				config.DbName))
	}

	if config.DbUser != expected.DbUser {
		t.Error(
			fmt.Sprintf("DbUser does not match: expected = %s, actual %s",
				expected.DbUser,
				config.DbUser))
	}

	if config.DbPassword != expected.DbPassword {
		t.Error(
			fmt.Sprintf("DbPassword does not match: expected = %s, actual %s",
				expected.DbPassword,
				config.DbPassword))
	}

}

func TestEnv(t *testing.T) {

	expected := HexGameOptions{
		HexGameConfig: HexGameConfig{
			ConfigFile: "a file name",
			ContentRoot: "where to put it",
			MongoDBConfig: MongoDBConfig{
				DbServer: "aserver",
				DbPort: 12345,
				DbName: "database",
				DbUser: "menotyou",
				DbPassword: "ssshdonttell",
			},
		},
		Verbose: true,
		Debug: false,
	}
	
	os.Setenv("HEXGAME_CONFIG_FILE", expected.ConfigFile)
	os.Setenv("HEXGAME_CONTENT_ROOT", expected.ContentRoot)
	
	os.Setenv("HEXGAME_DBSERVER", expected.DbServer)
	os.Setenv("HEXGAME_DBPORT", string(expected.DbPort))
	os.Setenv("HEXGAME_DBNAME", expected.DbName)
	os.Setenv("HEXGAME_DBUSER", expected.DbUser)
	os.Setenv("HEXGAME_DBPASS", expected.DbPassword)
	
// 	// prepare for the test

	actual := processEnv()

	if actual.ConfigFile != expected.ConfigFile {
		t.Error(
			fmt.Sprintf(
				"wrong config file string: expected: %s, actual %s",
				expected.ConfigFile,
				actual.ConfigFile))
	}

	if actual.ContentRoot != expected.ContentRoot {
		t.Error(
			fmt.Sprintf(
				"wrong content root string: expected: %s, actual %s",
				expected.ContentRoot,
				actual.ContentRoot))
	}

}
