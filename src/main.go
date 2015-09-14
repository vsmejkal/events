package main

import (
    "fmt"
	"model"
    "time"
)

func main() {
    query := model.EventQuery{
        From: model.Datetime{time.Now()},
    }

    for event := range query.Search() {
	   fmt.Println(event.Name) 
    }

    if err := query.Error(); err != nil {
        fmt.Println(err)
    }
}
