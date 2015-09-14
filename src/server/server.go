package main

import (
    "fmt"
    "net/http"
    "log"
    "io"
    "model"
)


func handleRoot(w http.ResponseWriter, r *http.Request) {
    events := model.GetEvents()

    for event := range events {
        fmt.
    }
}

func main() {
    err := model.Connect()
    if err != nil {
        fmt.Fatalf("Cannot establish db connection: %s", err)
    }
    defer model.Disconnect()

    fmt.Println("Listening on port 8000...")

    http.HandleFunc("/", handleRoot)
    err = http.ListenAndServe(":8000", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}