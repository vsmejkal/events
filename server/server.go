package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"time"
	"github.com/vsmejkal/events/model"
)


func getEventsQuery() model.EventQuery {
	// From today
	year, month, day := time.Now().Date()
	from := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	// To one week ahead
	to := time.Now().Add(7 * 24 * time.Hour)

	return model.EventQuery{
		From: model.Datetime{from},
		To: model.Datetime{to},
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	query := getEventsQuery()
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