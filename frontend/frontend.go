package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"time"
	"github.com/vsmejkal/events/model"
	"os"
	"github.com/vsmejkal/events/config"
	"path"
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
	tpl, err := template.ParseFiles(config.Frontend.DocumentRoot + "views/index.html")
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

func printUsage() {
	fmt.Printf("Usage: %s config.json\n", path.Base(os.Args[0]))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		printUsage()
		return
	}

	if err := config.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	if err := model.Connect(); err != nil {
		log.Fatal("Cannot establish db connection: %s", err)
	}
	defer model.Disconnect()

	http.HandleFunc("/", handleRoot)

	port := config.Frontend.Port
	fmt.Printf("Listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}