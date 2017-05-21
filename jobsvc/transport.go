package jobsvc

import (
	"database/sql"

	"github.com/julienschmidt/httprouter"
)

var svc Service
var endpoint Endpoint

func Register(router *httprouter.Router, database *sql.DB) {
	svc = Service{
		DB: database,
	}
	router.GET("/jobs", endpoint.GetJobs(svc))
	router.GET("/jobs/:id", endpoint.GetJob(svc))
	router.POST("/jobs", endpoint.CreateJob(svc))
	router.DELETE("/jobs/:id", endpoint.DeleteJob(svc))
	router.PUT("/jobs/:id", endpoint.UpdateJob(svc))
}
