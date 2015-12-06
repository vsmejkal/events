package main

import (
	"fmt"
	"time"
	"github.com/vsmejkal/events/model"
	"os"
	"github.com/vsmejkal/events/config"
	"log"
	"path"
)

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

	query := model.EventQuery{
		From: model.Datetime{time.Now()},
		// From: model.Datetime{time.Date(2015, time.September, 26, 17, 0, 0, 0, time.UTC)},
	}

	for event := range query.Search() {
		fmt.Println(event.Start, event.Name)
	}

	if err := query.Error(); err != nil {
		fmt.Println(err)
	}
}
