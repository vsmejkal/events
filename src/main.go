package main

import (
	"fmt"
	"log"
	"model"
	"parser"
)

func main() {
	eventChan := make(chan model.Event, 100)
	errChan := make(chan error, 100)

	go parser.ParseEvents("https://www.facebook.com/TwoFacesClub", eventChan, errChan)

	select {
	case event := <-eventChan:
		fmt.Println(event.Name)
	case err := <-errChan:
		log.Println(err)
	}
}
