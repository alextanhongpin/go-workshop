package jobsvc

// Job is the schema for the job model in the database
type Job struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
