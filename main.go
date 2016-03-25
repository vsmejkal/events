package main

import (
	"fmt"
	"time"
	"github.com/vsmejkal/events/model"
)

func main() {
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
