package main

import (
    "fmt"
    "net/http"
    "html/template"
    "log"
    "time"
    "model"
)


func handleRoot(w http.ResponseWriter, r *http.Request) {
    tpl, err := template.ParseFiles("views/index.html")
    if err != nil {
        log.Fatalln(err)
    }

    query := model.EventQuery{ From: model.Datetime{ time.Now() } }
    events := query.Search()
 
    err = tpl.Execute(w, events)
    if err != nil {
        log.Println(err)
    }

    if err := query.Error(); err != nil {
        log.Println("EventQuery error:", err)
    }
}

func main() {
    err := model.Connect()
    if err != nil {
        log.Fatal("Cannot establish db connection: %s", err)
    }
    defer model.Disconnect()

    fmt.Println("Listening on port 8080...")

    http.HandleFunc("/", handleRoot)
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}