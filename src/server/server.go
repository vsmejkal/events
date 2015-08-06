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
    fmt.Println("Conn3", model.GetConnection())
}

func main() {
    var err error

    err = model.Connect()
    if err != nil {
        fmt.Errorf("Cannot establish db connection: ", err)
        return
    }

    fmt.Println("Listening on port 8000...")

    http.HandleFunc("/", handleRoot)
    err = http.ListenAndServe(":8000", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}