package jobsvc

import (
	"fmt"
	"strings"
)

// Mock jobs in telecommunication companies
var jobs = []Job{
	{1, "Axiata Group Berhad"},
	{2, "Digi.com Berhad"},
	{3, "Maxis Berhad"},
	{4, "Telekom Malaysia Berhad"},
}

type Service struct{}

type getJobsRequest struct {
	Query string `json:"query"`
}

type getJobsResponse struct {
	Data []Job `json:"data"`
}

// GetJobs returns a list of jobs
func (svc Service) GetJobs(request getJobsRequest) (getJobsResponse, error) {
	query := request.Query

	var filteredJobs []Job

	// We only filter the jobs that fits the query
	for _, value := range jobs {
		if strings.Contains(strings.ToLower(value.Name), query) {
			filteredJobs = append(filteredJobs, value)
		}
	}
	return getJobsResponse{Data: filteredJobs}, nil
}

type getJobRequest struct {
	ID int `json:"id"`
}

type getJobResponse struct {
	*Job
}

// GetJob service returns a list of jobs
func (svc Service) GetJob(request getJobRequest) (getJobResponse, error) {
	id := request.ID

	var job *Job
	for _, value := range jobs {
		if value.ID == id {
			job = &value
			break
		}
	}
	if job == nil {
		return getJobResponse{}, fmt.Errorf("The job with the id %d is not found", id)

	}
	return getJobResponse{job}, nil
}
