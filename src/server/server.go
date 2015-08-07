package main

import (
    "fmt"
    "net/http"
    "log"
    "io"
    "model"
)


func handleRoot(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello!")
}

func main() {
    var err error
    err = model.Connect()
    defer model.Disconnect()
    if err != nil {
        fmt.Errorf("Cannot establish db connection: %s", err)
        return
    }

    db := model.GetConnection()
    result, err := db.Exec("INSERT INTO event(name, link, starttime, endtime) VALUES($1,$2,$3,$4);", "XXX", "YYY", "2012-12-09", "2012-12-09")
    if (err != nil) {
        fmt.Println(err)
        return
    }

    fmt.Println("Listening on port 8000...")

    http.HandleFunc("/", handleRoot)
    err = http.ListenAndServe(":8000", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}