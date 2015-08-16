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

func SerializeTime(tm time.Time) string {
    return tm.Format(time.RFC3339)
}

func DeserializeTime(data string) time.Time {
    tm, _ := time.Parse(time.RFC3339, data)

    return tm
}

func SerializeStringArray(arr []string) string {
    return "{" + strings.Join(arr, ",") + "}"
}

func DeserializeStringArray(data string) []string {
    return strings.Split(data[1:len(data)-1], ",")
}

func SerializeIntArray(arr []int) string {
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

func DeserializeIntArray(data string) []int {
    strs := strings.Split(data, ",")
    ints := make([]int, len(strs), len(strs))

    for i, n := range strs {
        ints[i], _ = strconv.Atoi(n)
    }

    return ints
}