package app

// config.go contains the implementation to setup the app config by reading the config from a json file

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strconv"
)

// Configuration will contain the config for our program
type configuration struct {
	// The port for the server. Defaults to 8080.
	Port int `json:"port"`
	// The data source name (DSN) to connect to the database. Required.
	DSN string `json:"dsn"`
}

// GetPort returns the int port to ":8080" format
func (cfg configuration) GetPort() string {
	// Convert the port to string
	return ":" + strconv.Itoa(cfg.Port)
}

var Config configuration

var (
	cfg  = flag.String("config", "env/dev.json", "The path to the config file.")
	port = flag.Int("port", 8080, "The port of the server")
)

// SetupConfig will initialize the config
func SetupConfig() {
	flag.Parse()

	f, _ := os.Open(*cfg)
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&Config); err != nil {
		log.Fatal(err)
	}

	Config.Port = *port
}
