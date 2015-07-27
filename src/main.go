package main

import (
    "fmt"
    "log"
    "parser"
    "model"
)

func main() {
    eventChan := make(chan model.Event, 100)
    errChan := make(chan error, 100)

    go parser.ParseEvents("https://www.facebook.com/TwoFacesClub", eventChan, errChan)

    select {
    case event := <-eventChan:
        fmt.Println(event.Start)
    case err := <-errChan:
        log.Println(err)
    }
}
