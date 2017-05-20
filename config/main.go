package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Configuration will contain the config for our program
type Configuration struct {
	Port int `json:"port"`
}

// Define a variable that belongs to the struct `Configuration`
var configuration Configuration

// init will be called before the program executes
func init() {
	loadConfig(&configuration)
}

func loadConfig(configuration *Configuration) {
	// Read the ENV from the command line
	// $ ENV=DEVELOPMENT go run main.go
	env := os.Getenv("ENV")
	if env == "" {
		fmt.Println("ConfigError: The ENV is not defined. Defaults to `DEVELOPMENT`.")
		env = "DEVELOPMENT"
	}
	pathToConfig := checkEnvironment(env)
	file, _ := os.Open(pathToConfig)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("config.go: Loaded config for environment " + strings.ToLower(env) + ".")
}

// checkEnvironment will return the path to config based on the env
func checkEnvironment(env string) string {
	switch env {
	case "DEVELOPMENT":
		return "env/dev.json"
	case "STAGING":
		return "env/stage.json"
	case "PRODUCTION":
		return "env/prod.json"
	case "TESTING":
		return "env/test.json"
	default:
		return "env/dev.json"
	}
}

// Read will return the config that is loaded
func Read() Configuration {
	return configuration
}
