package server

import (
	"testing"

	"fmt"
	"os"
	"strconv"
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
	os.Setenv("HEXGAME_DBPORT", strconv.Itoa(expected.DbPort))
	os.Setenv("HEXGAME_DBNAME", expected.DbName)
	os.Setenv("HEXGAME_DBUSER", expected.DbUser)
	os.Setenv("HEXGAME_DBPASS", expected.DbPassword)

	os.Setenv("HEXGAME_VERBOSE", strconv.FormatBool(expected.Verbose))
	os.Setenv("HEXGAME_DEBUG", strconv.FormatBool(expected.Debug))
	
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

	if actual.DbServer != expected.DbServer {
		t.Error(
			fmt.Sprintf(
				"wrong db server: expected: %s, actual %s",
				expected.DbServer,
				actual.DbServer))
	}

	if actual.DbPort != expected.DbPort {
		t.Error(
			fmt.Sprintf(
				"wrong db port: expected: %d, actual %d",
				expected.DbPort,
				actual.DbPort))
	}

	if actual.DbName != expected.DbName {
		t.Error(
			fmt.Sprintf(
				"wrong db name: expected: %s, actual %s",
				expected.DbName,
				actual.DbName))
	}

	if actual.DbUser != expected.DbUser {
		t.Error(
			fmt.Sprintf(
				"wrong db user: expected: %s, actual %s",
				expected.DbUser,
				actual.DbUser))
	}

	if actual.DbPassword != expected.DbPassword {
		t.Error(
			fmt.Sprintf(
				"wrong db password: expected: %s, actual %s",
				expected.DbPassword,
				actual.DbPassword))
	}


	if actual.Verbose != expected.Verbose {
		t.Error(
			fmt.Sprintf(
				"wrong verbose flag: expected: %v, actual %v",
				expected.Verbose,
				actual.Verbose))
	}
	
	if actual.Debug != expected.Debug {
		t.Error(
			fmt.Sprintf(
				"wrong debug flag: expected: %v, actual %v",
				expected.Debug,
				actual.Debug))
	}

}

func TestMergeConfig(t *testing.T) {

	empty := HexGameConfig{}

	// var full :=	HexGameConfig: HexGameConfig{
	// 	ConfigFile: "a file name",
	// 	ContentRoot: "where to put it",
	// 	MongoDBConfig: MongoDBConfig{
	// 		DbServer: "aserver",
	// 		DbPort: 12345,
	// 		DbName: "database",
	// 		DbUser: "menotyou",
	// 		DbPassword: "ssshdonttell",
	// 	},
	// 	Verbose: true,
	// 	Debug: false,
	//}

	input := HexGameConfig{
		ConfigFile: "configfile",
		ContentRoot: "location",
		MongoDBConfig: MongoDBConfig{
			DbServer: "servername",
			DbPort: 23456,
			DbName: "dbname",
			DbUser: "dbuser",
			DbPassword: "password",
		},
	}

	output := MergeConfigs(&empty, &empty)

	if output.ContentRoot != empty.ContentRoot {
		t.Error(fmt.Sprintf("empty merge failed: expected: %s, actual: %s",
			empty.ContentRoot,
			output.ContentRoot))
	}

	output = MergeConfigs(&input, &empty)

	if output.ContentRoot != input.ContentRoot {
		t.Error(fmt.Sprintf("merge failed: expected: %s, actual: %s",
			input.ContentRoot,
			output.ContentRoot))
	}

	output = MergeConfigs(&empty, &input)

	if output.ContentRoot != input.ContentRoot {
		t.Error(fmt.Sprintf("merge failed: expected: %s, actual: %s",
			input.ContentRoot,
			output.ContentRoot))
	}

}
