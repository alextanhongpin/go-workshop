package jobsvc

import (
	"github.com/alextanhongpin/go-workshop/database"

	"github.com/julienschmidt/httprouter"
)

var service Service
var endpoint Endpoint

func init() {
	endpoint = Endpoint{}
	service = Service{
		DB: database.DB(),
	}
}

func Register(router *httprouter.Router) {

	router.GET("/jobs", endpoint.GetJobs(service))
	router.GET("/jobs/:id", endpoint.GetJob(service))
	router.POST("/jobs", endpoint.CreateJob(service))
	router.DELETE("/jobs/:id", endpoint.DeleteJob(service))
	router.PUT("/jobs/:id", endpoint.UpdateJob(service))
}
