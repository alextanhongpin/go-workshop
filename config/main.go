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
	env := os.Getenv("env")
	if env == "" {
		fmt.Println("ConfigInfo: The `env` flag is not defined. Defaults to `development`.")
		env = "development"
	}
	path := getPath(env)
	readFile(path, &configuration)
}

func readFile(path string, configuration *Configuration) {
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

// getPath will return the path to config based on the env
func getPath(env string) string {
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
