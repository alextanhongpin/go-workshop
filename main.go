package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/alextanhongpin/go-workshop/config"
	"github.com/alextanhongpin/go-workshop/jobsvc"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Services interface {
	Register(router *httprouter.Router)
}

func main() {
	config := config.Read()
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		p := User{
			Name: "Hello",
			Age:  1,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
		// w.Write(out)
	})
	jobsvc.Register(router)
	fmt.Println("listening to port *:" + strconv.Itoa(config.Port) + ". press ctrl + c to cancel.")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router))
}
