package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alextanhongpin/go-workshop/app"
	"github.com/alextanhongpin/go-workshop/jobsvc"

	"github.com/julienschmidt/httprouter"
)

// version is the current version of the server
const version string = "0.0.1"

func main() {
	// Read and load the config for our app
	app.SetupConfig()

	// Setup our mysql database with the given Data Source Name (DNS)
	app.SetupDatabase(app.Config.DSN)

	// Setup router
	router := httprouter.New()

	// Register the job-service. Optionally you can add a feature toggle here.
	jobsvc.Register(router, app.DB)

	fmt.Printf("listening to port *%s. press ctrl + c to cancel.", app.Config.GetPort())
	log.Fatal(http.ListenAndServe(app.Config.GetPort(), router))
}
