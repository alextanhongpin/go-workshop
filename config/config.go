package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration will contain the config for our program
type Configuration struct {
	Port       int    `json:"port"`
	DBHost     string `json:"db_host"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBDatabase string `json:"db_database"`
}

const configPath = "CONFIG"

var path string = os.Getenv(configPath)
var configuration Configuration

// init will be called before the program executes
func init() {
	// Path to config file
	// path := os.Getenv(configPath)
	// unset apiEnv to avoid side-effect for future env and flag parsing.
	os.Unsetenv(configPath)

	if path == "" {
		path = "env/dev.json"
	}

	loadConfig(path, &configuration)
}

func loadConfig(path string, configuration *Configuration) {

	file, _ := os.Open(path)
	defer file.Close()
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
