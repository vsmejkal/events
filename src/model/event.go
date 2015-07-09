package model

import "time"

type Event struct {
    Id uint64
	Name string
	Desc string
	Link string
	Image string
    Start time.Time
    End time.Time
    IsDateOnly bool
    Place Place
    Categories []string
}

func (e *Event) IsValid() bool {
    return e.Name != "" &&
           e.Link != "" &&
           e.Start.After(time.Now()) &&
           e.End.After(time.Now()) &&
           e.Place.IsValid()
}