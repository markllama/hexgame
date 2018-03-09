package server

import (
	"testing"
)

func TestLoad(t *testing.T) {

	test_config_file := "../test/hexgame_config.json"

	LoadConfig(test_config_file)

	t.Fail()
}
