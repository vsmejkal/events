package model

import "time"

type Event struct {
	Id       uint64
	Name     string
	Desc     string
	Link     string
	Image    string
	Start    time.Time
	End      time.Time
	DateOnly bool
	Place    Place
	Tags     []string
}

func FindEvents() *EventQuery {
	// SQL query
	return &EventQuery{}
}

func (e *Event) Store() {
	if !e.IsDuplicate() {
		// SQL query
	}
}

func (e *Event) IsDuplicate() bool {
	// SQL query
	return false
}

func (e *Event) IsValid() bool {
	return e.Name != "" &&
		e.Link != "" &&
		e.Start.After(time.Now()) &&
		e.End.After(time.Now()) &&
		e.Place.IsValid()
}
