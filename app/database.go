package app

// database.go contains the implementation for our app database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	driverName   string = "mysql"
	maxOpenConns int    = 150
)

func SetupDatabase(dataSourceName string) {
	var err error
	DB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetMaxOpenConns(maxOpenConns)

	// Check if the job table exists and create one if not
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS job (
		id INT NOT NULL AUTO_INCREMENT,
		name TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	)`)

	if err != nil {
		log.Fatal(err)
	}

	// Ping to ensure connection is available
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
