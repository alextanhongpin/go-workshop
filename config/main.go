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

// init will be called before the program executes
func init() {
	loadConfig(&configuration)
}

func loadConfig(configuration *Configuration) {
	// Read the ENV from the command line
	// $ ENV=DEVELOPMENT go run main.go
	env := os.Getenv("env")
	if env == "" {
		fmt.Println("ConfigError: The `env` flag is not defined. Defaults to `development`.")
		env = "development"
	}
	pathToConfig := checkEnvironment(env)
	file, _ := os.Open(pathToConfig)
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&configuration); err != nil {
		fmt.Println("ConfigError: ", err)
	}
	fmt.Println("config.go: Loaded config for environment " + env + ".")
}

// checkEnvironment will return the path to config based on the env
func checkEnvironment(env string) string {
	switch env {
	case "development":
		return "env/dev.json"
	case "staging":
		return "env/stage.json"
	case "production":
		return "env/prod.json"
	case "testing":
		return "env/test.json"
	default:
		return "env/dev.json"
	}
}

// Read will return the config that is loaded
func Read() Configuration {
	return configuration
}
