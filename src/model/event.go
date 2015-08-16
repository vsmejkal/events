package model

import (
	"fmt"
	"time"
)

type Event struct {
	Id       int64
	Name     string
	Desc     string
	Link     string
	Image    string
	Start    time.Time
	End      time.Time
	Tags     []string
	Place    Place
}

func FindEvents() *EventQuery {
	// SQL query
	return &EventQuery{}
}

func (e *Event) Store() error {
	if err := e.Place.Store(); err != nil {
		return err
	}

	db := GetConnection()

	err := db.QueryRow("INSERT INTO event(name, description, link, image, starttime, endtime, tags, place) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id;",
				e.Name,
				e.Desc,
				e.Link,
				e.Image,
				SerializeTime(e.Start),
				SerializeTime(e.End),
				SerializeStringArray(e.Tags),
				e.Place.Id,
			).Scan(&e.Id);
    
    if err != nil {
        return fmt.Errorf("Event.Store() error: %s", err)
    }

    return nil
}

func (e *Event) IsDuplicate() bool {
	// db := GetConnection()
	return false
}

func (e *Event) IsValid() bool {
	return e.Name != "" &&
		e.Link != "" &&
		e.Start.After(time.Now()) &&
		e.Place.IsValid()
}