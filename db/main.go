package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Define the MySQL database connection parameters
    username := "root"
    password := "password"
    host := "localhost"
    port := "3306"
    database := "tharun"

    // Create a DSN (Data Source Name) string
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

    // Open a connection to the MySQL database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // Defer closing the database connection until the function exits

    // Attempt to ping the database to check the connection
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Connected to the database!")

    // Now you can perform database operations using the 'db' object
    // For example, you can execute queries, insert data, etc.
    // ...

}

