package jobsvc

import (
	// Standard packages
	"encoding/json"
	"net/http"
	"strconv"

	// Your packages
	"github.com/alextanhongpin/go-workshop/httputil"

	// Vendor packages
	"github.com/julienschmidt/httprouter"
)

type Endpoint struct{}

// GetJobs will return a list of jobs
func (Endpoint) GetJobs(service Service) httprouter.Handle {
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
		httputil.Json(w, res, http.StatusOK)
	}
}

// GetJob will return a job by id
func (Endpoint) GetJob(service Service) httprouter.Handle {
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
		httputil.Json(w, res, http.StatusOK)
	}
}

func (Endpoint) CreateJob(service Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		var req createJobRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := service.CreateJob(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httputil.Json(w, res, http.StatusCreated)
	}
}

func (Endpoint) DeleteJob(service Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		var req deleteJobRequest
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
		req.ID = id

		res, err := service.DeleteJob(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httputil.Json(w, res, http.StatusNoContent)
	}
}

func (Endpoint) UpdateJob(service Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		var req updateJobRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
		req.ID = id

		res, err := service.UpdateJob(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httputil.Json(w, res, http.StatusNoContent)
	}
}
