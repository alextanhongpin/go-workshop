package jobsvc

import (
	"time"
)

// Job is the schema for the job model in the database
type Job struct {
	ID        int        `db:"id" json:"id,omitempty"`
	Name      string     `db:"name" json:"name,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"` // We use a pointer here so that it will be nil
}

// http://www.zbeanztech.com/blog/important-mysql-commands
