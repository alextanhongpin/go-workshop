package jobsvc

import (
	"time"
)

// Job is the schema for the job model in the database
type Job struct {
	// ID is a the unique primary key for the job
	ID int `db:"id" json:"id,omitempty"`
	// Name is the name of the job entry
	Name string `db:"name" json:"name,omitempty"`
	// CreatedAt is the time when the entry is created
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"` // We use a pointer here so that it will be nil
}
