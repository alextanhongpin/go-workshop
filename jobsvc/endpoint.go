package jobsvc

import (
	// Standard packages
	"net/http"
	"strconv"

	// Your packages
	"github.com/alextanhongpin/go-workshop/httputil"

	// Vendor packages
	"github.com/julienschmidt/httprouter"
)

type Endpoint struct{}

// GetJobs will return a list of jobs
func (endpoint Endpoint) GetJobs(service Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Construct a request
		req := getJobsRequest{
			Query: r.URL.Query().Get("query"), //r.FormValue("query"),
		}

		// Call the service with the request
		res, err := service.GetJobs(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Return the payload as json
		httputil.Json(w, res)
	}
}

// GetJob will return a job by id
func (endpoint Endpoint) GetJob(service Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the string id and convert it to int
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			httputil.Error(w, "The id provided is malformed", http.StatusBadRequest)
			return
		}

		// Construct a request
		req := getJobRequest{
			ID: id,
		}

		// Call the service with the request
		res, err := service.GetJob(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httputil.Json(w, res)
	}
}
