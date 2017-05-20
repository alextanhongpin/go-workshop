package jobsvc

import (
	"github.com/julienschmidt/httprouter"
)

func Register(router *httprouter.Router) {
	var service Service
	service = Service{}

	var endpoint Endpoint
	endpoint = Endpoint{}

	router.GET("/jobs", endpoint.GetJobs(service))
	router.GET("/jobs/:id", endpoint.GetJob(service))
}
