package jobsvc

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
	DB *sql.DB
}

type getJobsRequest struct {
	Query string `json:"query"`
}

type getJobsResponse struct {
	Data []Job `json:"data"`
}

// GetJobs returns a list of jobs
func (svc Service) GetJobs(request getJobsRequest) (getJobsResponse, error) {
	// query := request.Query

	var response []Job
	// Prepare a statement
	stmt, err := svc.DB.Prepare("SELECT id, name, created_at FROM job")
	defer stmt.Close()
	if err != nil {
		log.Fatal("[PrepareError]", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("[QueryError]", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var createdAt *time.Time

		err = rows.Scan(&id, &name, &createdAt)

		if err != nil {
			log.Fatal("[ScanError]", err)
		}

		response = append(response, Job{
			ID:        id,
			Name:      name,
			CreatedAt: createdAt,
		})
	}
	// rows.Err() will throw an error if the operation in rows.Next() fails
	if err = rows.Err(); err != nil {
		fmt.Println("[RowsError]", err)
		return getJobsResponse{}, err
	}
	return getJobsResponse{Data: response}, nil
}

type getJobRequest struct {
	ID int `json:"id"`
}

type getJobResponse struct {
	Data Job `json:"data,omitempty"`
}

// GetJob service returns a list of jobs
func (svc Service) GetJob(request getJobRequest) (getJobResponse, error) {
	id := request.ID
	fmt.Println("GetJobRequest:", id)

	var job Job

	stmt, err := svc.DB.Prepare("SELECT id, name, created_at FROM job WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(id).Scan(&job.ID, &job.Name, &job.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// There are no rows, but otherwise no errors occured
			return getJobResponse{}, nil
		} else {
			log.Fatal(err)
		}
	}
	return getJobResponse{job}, nil
}

type createJobRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type createJobResponse struct {
	ID int64 `json:"id"`
}

func (svc Service) CreateJob(request createJobRequest) (createJobResponse, error) {
	stmt, err := svc.DB.Prepare("INSERT INTO job (name) values (?)")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	// Use exec for UPDATE, DELETE, POST operation
	res, err := stmt.Exec(request.Name)
	if err != nil {
		fmt.Println(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}

	return createJobResponse{ID: id}, nil
}

type deleteJobRequest struct {
	ID int64 `json:"id"`
}
type deleteJobResponse struct {
	Ok bool  `json:"ok"` // The success message
	N  int64 `json:"n"`  // The number of rows deleted
}

func (svc Service) DeleteJob(request deleteJobRequest) (deleteJobResponse, error) {
	stmt, err := svc.DB.Prepare("DELETE FROM job WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return deleteJobResponse{}, err
	}

	res, err := stmt.Exec(request.ID)
	if err != nil {
		return deleteJobResponse{}, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return deleteJobResponse{}, err
	}
	return deleteJobResponse{Ok: true, N: n}, nil
}

type updateJobRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type updateJobResponse struct {
	Ok bool  `json:"ok"`
	N  int64 `json:"n"`
}

func (svc Service) UpdateJob(request updateJobRequest) (updateJobResponse, error) {
	stmt, err := svc.DB.Prepare("UPDATE job SET name=? WHERE id=?")
	defer stmt.Close()

	if err != nil {
		return updateJobResponse{}, nil
	}

	res, err := stmt.Exec(request.Name, request.ID)
	if err != nil {
		return updateJobResponse{}, nil
	}
	n, err := res.RowsAffected()
	if err != nil {
		return updateJobResponse{}, nil
	}
	return updateJobResponse{Ok: true, N: n}, nil
}

// https://dev.mysql.com/downloads/mysql/
// https://askubuntu.com/questions/408676/accessing-mysql-using-terminal-in-ubuntu-13-04
// 2017-05-20T18:56:04.204985Z 1 [Note] A temporary password is generated for root@localhost: Coo:4(C=l0wo

// If you lose this password, please consult the section How to Reset the Root Password in the MySQL reference manual.
// mysql -u root -p

// Difference between prepare and query
//http://stackoverflow.com/questions/37404989/whats-the-difference-between-db-query-and-db-preparestmt-query-in-golang
