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
		{"STAGING", "env/stage.json"},
		{"DEVELOPMENT", "env/dev.json"},
		{"PRODUCTION", "env/prod.json"},
		{"TESTING", "env/test.json"},
	}

	for _, value := range environments {
		os.Setenv("ENV", value.env)
		env := os.Getenv("ENV")
		actual := checkEnvironment(env)
		expected := value.path
		if actual != expected {
			t.Errorf("Test failed, expected %s, got %s", expected, actual)
		}
	}
}
