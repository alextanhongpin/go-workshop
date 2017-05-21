package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alextanhongpin/go-workshop/config"

	_ "github.com/go-sql-driver/mysql"
)

// Create a database first
// db, err := sql.Open("mysql", "root:123456@gotest")

var db *sql.DB

// Init will be call before the package executes
func init() {
	var err error
	config := config.Read()
	driverName := "mysql"

	// Construct the driver-specific syntax to access the datastore

	// Faster than concatenating strings
	// var syntax bytes.Buffer
	// syntax.WriteString(config.DBUser)
	// syntax.WriteString(":")
	// syntax.WriteString(config.DBPassword)
	// syntax.WriteString("@/")
	// syntax.WriteString(config.DBDatabase)
	// syntax.String()

	// You need to set parseTime
	driverPath := fmt.Sprintf("%s:%s@/?parseTime=true", config.DBUser, config.DBPassword)
	db, err = sql.Open(driverName, driverPath)
	if err != nil {
		log.Fatal(err)
	}

	// Only create the DB if it does not exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DBDatabase)
	if err != nil {
		log.Fatal(err)
	}

	// Error 1046: No database selected

	// Use the database with the name provided from config
	_, err = db.Exec("USE " + config.DBDatabase)
	if err != nil {
		log.Fatal(err)
	}

	// Create a tables that we are using
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS job (
		id INT NOT NULL AUTO_INCREMENT,
		name TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`)

	if err != nil {
		log.Fatal(err)
	}
	// Ping to ensure connection is available
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	// Do not close connections unnecessarily
	// defer db.Close()
}

// DB returns a reference to the db
func DB() *sql.DB {
	return db
}

// Go-database-sql overview
// http://go-database-sql.org/overview.html
// curl -X POST -d '{"name":"test"}' http://localhost:8080/jobs
