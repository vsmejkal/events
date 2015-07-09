package main

import (
    "fmt"
    "log"
    "parser"
    "model"
)

func main() {
    events := make([]model.Event, 0, 20)
    err := parser.ParseEvents("https://www.facebook.com/TwoFacesClub", &events)

    if (err != nil) {
        log.Fatal(err)
        return
    }

    for _, v := range events {
        fmt.Println(v.Start)
    }
}
