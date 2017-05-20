package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration will contain the config for our program
type Configuration struct {
	Port int `json:"port"`
}

// Define a variable that belongs to the struct `Configuration`
var configuration Configuration

const configPath = "CONFIG"

// init will be called before the program executes
func init() {
	// Path to config file
	path := os.Getenv(configPath)
	// unset apiEnv to avoid side-effect for future env and flag parsing.
	os.Unsetenv(configPath)

	loadConfig(path, &configuration)
}

func loadConfig(path string, configuration *Configuration) {
	if path == "" {
		fmt.Println("ConfigError: path to config is not defined")
		return
	}
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configuration); err != nil {
		fmt.Println("ConfigError:", err)
		return
	}
}

// Read will return the config that is loaded
func Read() Configuration {
	return configuration
}
