package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/alextanhongpin/go-workshop/config"
	"github.com/alextanhongpin/go-workshop/jobsvc"

	"github.com/julienschmidt/httprouter"
)

type Services interface {
	Register(router *httprouter.Router)
}

func main() {
	config := config.Read()
	router := httprouter.New()

	jobsvc.Register(router)
	fmt.Printf("listening to port *:%d. press ctrl + c to cancel.", config.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router))
}
