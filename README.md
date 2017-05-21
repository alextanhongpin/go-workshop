# Go Workshop

Structuring a golang's web api


## Useful commands

Go imports
```bash
$ go get golang.org/x/tools/cmd/goimports
$ goimport
```


## Benchmarking HTTP

```bash

$ brew install wrk
$ wrk -c100 -d10 -t10 "http://localhost:8080/jobs"
```

// Faster than concatenating strings
// var syntax bytes.Buffer
// syntax.WriteString(config.DBUser)
// syntax.WriteString(":")
// syntax.WriteString(config.DBPassword)
// syntax.WriteString("@/")
// syntax.WriteString(config.DBDatabase)
// syntax.String()

    // Create a database if it doesn't exist
    _, err = DB.Exec("CREATE DATABASE IF NOT EXISTS " + config.DBDatabase)
    if err != nil {
        log.Fatal(err)
    }

    // Error 1046: No database selected

    // Use the database with the name provided from config
    _, err = DB.Exec("USE " + config.DBDatabase)
    if err != nil {
        log.Fatal(err)
    }

    // Create a tables that we are using
    _, err = DB.Exec(`CREATE TABLE IF NOT EXISTS job (
        id INT NOT NULL AUTO_INCREMENT,
        name TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`)

// http://www.zbeanztech.com/blog/important-mysql-commands


// https://dev.mysql.com/downloads/mysql/
// https://askubuntu.com/questions/408676/accessing-mysql-using-terminal-in-ubuntu-13-04
// 2017-05-20T18:56:04.204985Z 1 [Note] A temporary password is generated for root@localhost: Coo:4(C=l0wo

// If you lose this password, please consult the section How to Reset the Root Password in the MySQL reference manual.
// mysql -u root -p

// Difference between prepare and query
//http://stackoverflow.com/questions/37404989/whats-the-difference-between-db-query-and-db-preparestmt-query-in-golang
