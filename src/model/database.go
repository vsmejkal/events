package model

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

const (
    DB_NAME = "eventsdb"
)

// Database handle
var conn *sql.DB

func Connect() (err error) {
    dbinfo := fmt.Sprintf("dbname=%s", DB_NAME)

    conn, err = sql.Open("postgres", dbinfo)
    if err != nil {
        log.Fatal("Database error: ", err)
    }

    err = conn.Ping()
    if err != nil {
        log.Fatal("Database error: ", err)
    }

    return err
}

func Disconnect() {
    if conn != nil {
        conn.Close()
    }
}

func GetConnection() *sql.DB {
    if conn == nil {
        Connect()
    }

    return conn
}