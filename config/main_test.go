package config

import (
	"os"
	"testing"
)

// TestCheckEnvironment will validate if the configs are loaded correctly according to the environment set
func TestCheckEnvironments(t *testing.T) {
	environments := []struct {
		env  string
		path string
	}{
		{"staging", "env/stage.json"},
		{"development", "env/dev.json"},
		{"production", "env/prod.json"},
		{"testing", "env/test.json"},
	}

	for _, value := range environments {
		os.Setenv("env", value.env)
		env := os.Getenv("env")
		actual := checkEnvironment(env)
		expected := value.path
		if actual != expected {
			t.Errorf("Test failed, expected %s, got %s", expected, actual)
		}
	}
}
