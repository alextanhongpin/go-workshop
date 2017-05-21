package jobsvc

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type service interface {
	// all returns all the jobs available.
	All(allReq) (allRes, error)
	// one returns the job with the specified id.
	One(oneReq) (oneRes, error)
	// create saves a job in the database.
	Create(createReq) (createRes, error)
	// count returns the number of jobs.
	// Count(countReq) (countRes, error)
	// update updates the job with the given id in the database.
	Update(updateReq) (updateRes, error)
	// delete removes the job with the given id from the database.
	Delete(deleteReq) (deleteRes, error)
}

type Service struct {
	DB *sql.DB
}

type allReq struct{}

type allRes struct {
	Data []Job `json:"data"`
}

func (svc Service) All(req allReq) (allRes, error) {
	var err error
	var res []Job
	rows, err := svc.DB.Query("SELECT id, name, created_at FROM job")
	if err != nil {
		log.Fatal("[QueryError]", err)
	}
	defer rows.Close()
	for rows.Next() {
		var j Job

		err = rows.Scan(&j.ID, &j.Name, &j.CreatedAt)
		if err != nil {
			return allRes{}, err
		}

		res = append(res, Job{
			ID:        j.ID,
			Name:      j.Name,
			CreatedAt: j.CreatedAt,
		})
	}

	// rows.Err() will throw an error if the operation in rows.Next() fails
	if err = rows.Err(); err != nil {
		return allRes{}, err
	}
	return allRes{Data: res}, nil
}

type oneReq struct {
	ID int `json:"id"`
}

type oneRes struct {
	Data Job `json:"data,omitempty"`
}

// GetJob service returns a list of jobs
func (svc Service) One(req oneReq) (oneRes, error) {
	id := req.ID
	var job Job
	err := svc.DB.QueryRow("SELECT id, name, created_at FROM job WHERE id = ?", id).Scan(&job.ID, &job.Name, &job.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// There are no rows, but otherwise no errors occured
			return oneRes{}, nil
		} else {
			log.Fatal(err)
		}
	}
	return oneRes{job}, nil
}

type createReq struct {
	Name string `json:"name"`
}

type createRes struct {
	ID int64 `json:"id"`
}

func (svc Service) Create(req createReq) (createRes, error) {
	res, err := svc.DB.Exec("INSERT INTO job (name) values (?)", req.Name)
	if err != nil {
		return createRes{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return createRes{}, err
	}

	return createRes{ID: id}, nil
}

type deleteReq struct {
	ID int64 `json:"id"`
}
type deleteRes struct {
	Ok bool  `json:"ok"` // The success message
	N  int64 `json:"n"`  // The number of rows deleted
}

func (svc Service) Delete(req deleteReq) (deleteRes, error) {
	res, err := svc.DB.Exec("DELETE FROM job WHERE id=?", req.ID)
	if err != nil {
		return deleteRes{}, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return deleteRes{}, err
	}
	return deleteRes{Ok: true, N: n}, nil
}

type updateReq struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type updateRes struct {
	Ok bool  `json:"ok"`
	N  int64 `json:"n"`
}

func (svc Service) Update(req updateReq) (updateRes, error) {
	res, err := svc.DB.Exec("UPDATE job SET name=? WHERE id=?", req.Name, req.ID)
	if err != nil {
		return updateRes{}, nil
	}
	n, err := res.RowsAffected()
	if err != nil {
		return updateRes{}, nil
	}
	return updateRes{Ok: true, N: n}, nil
}
