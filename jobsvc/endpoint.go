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
func (Endpoint) GetJobs(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Construct a request
		req := allReq{}

		// Call the service with the request
		res, err := svc.All(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Return the payload as json
		httputil.Json(w, res, http.StatusOK)
	}
}

// GetJob will return a job by id
func (Endpoint) GetJob(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the string id and convert it to int
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			httputil.Error(w, "The id provided is malformed", http.StatusBadRequest)
			return
		}

		// Construct a request
		req := oneReq{
			ID: id,
		}

		// Call the service with the request
		res, err := svc.One(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httputil.Json(w, res, http.StatusOK)
	}
}

func (Endpoint) CreateJob(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		var req createReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := svc.Create(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httputil.Json(w, res, http.StatusCreated)
	}
}

func (Endpoint) DeleteJob(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		var req deleteReq
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
		req.ID = id

		res, err := svc.Delete(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httputil.Json(w, res, http.StatusNoContent)
	}
}

func (Endpoint) UpdateJob(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		var req updateReq
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

		res, err := svc.Update(req)
		if err != nil {
			httputil.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httputil.Json(w, res, http.StatusNoContent)
	}
}
