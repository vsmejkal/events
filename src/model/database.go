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

	"config"
)

// Database handle
var conn *sql.DB

func Connect() (err error) {
	dbinfo := fmt.Sprintf("dbname=%s", config.DATABASE_NAME)

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

func serializeTime(tm time.Time) string {
	return tm.Format(time.RFC3339)
}

func deserializeTime(data string) time.Time {
	tm, _ := time.Parse(time.RFC3339, data)
	
	return tm
}

func serializeStringArr(arr []string) string {
	return "{" + strings.Join(arr, ",") + "}"
}

func deserializeStringArr(data string) []string {
	return strings.Split(data[1:len(data)-1], ",")
}

func serializeIntArr(arr []int) string {
	var buf bytes.Buffer

	// Opening bracket
	buf.WriteString("{")

	for _, n := range arr  {
		buf.WriteString(strconv.Itoa(n))
		buf.WriteString(",")
	}

	// Remove trailing comma
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}

	// Closing bracket
	buf.WriteString("}")

	return buf.String()
}

func deserializeIntArr(data string) []int {
	strs := strings.Split(data, ",")
	arr := make([]int, len(strs), len(strs))

	for i, n := range strs {
		arr[i], _ = strconv.Atoi(n)
	}

	return arr
}