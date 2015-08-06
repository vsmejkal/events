package model

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

const (
    DB_USER = "postgres"
    DB_PASS = "postgres"
    DB_NAME = "pgtest"
)

// Database handle
var conn *sql.DB


func Connect() (err error) {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s", DB_USER, DB_PASS, DB_NAME)
    conn, err = sql.Open("postgres", dbinfo)

    fmt.Println(err)

    if err != nil {
        log.Fatal(err)
    }

    return err
}

func GetConnection() *sql.DB {
    if conn == nil {
        Connect()
    }

    return conn
}