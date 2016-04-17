package model

import (
	"fmt"
	"log"
	"time"
	"bytes"
	"strings"
	"strconv"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/vsmejkal/events/config"
)

// Database handle
var conn *sql.DB

func Connect() (err error) {
	dbinfo := fmt.Sprintf("dbname=%s", config.Database.Name)

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
