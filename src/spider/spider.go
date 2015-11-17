package main

import (
	"fmt"
	"log"
	"time"
	"model"
	"parser"
)

func main() {
	sources := []string {
		"https://www.facebook.com/fledaclub",
	}

	for _, url := range sources {
		eventChan := make(chan model.Event, 100)
		errChan := make(chan error, 100)

		fmt.Println("Parsing", url, "...")

		go func() {
			parser.ParseEvents(url, eventChan, errChan)
			close(eventChan)
		}()

		loop: for {
			select {
			case event, ok := <-eventChan:
				if !ok {
					break loop
				}

				if event.IsValid() {
					fmt.Println("NEW:", event.Name, event.Start)
					
					if err := event.Store(); err != nil {
						log.Println(err)
					}
				} else {
					fmt.Println("OLD:", event.Name, event.Start)
				}

			case err := <-errChan:
				log.Println(err)
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}
}