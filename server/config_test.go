package server

import (
	"testing"
	"fmt"
)

func TestLoad(t *testing.T) {

	test_config_file := "../test/hexgame_config.json"

	config := LoadConfig(test_config_file)

	expected := struct {
		DbServer string
		DbPort int
	}{
		DbServer: "localhost",
		DbPort: 27017,
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
}
