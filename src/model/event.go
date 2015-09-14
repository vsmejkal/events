package model

import (
	"fmt"
	"time"
	"database/sql"
)

type Event struct {
	Id       int64
	Name     string
	Desc     string
	Link     string
	Image    string
	Start    Datetime
	End      Datetime
	Tags     Tags
	Place    Place
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
				e.Start.Encode(),
				e.End.Encode(),
				e.Tags.Encode(),
				e.Place.Id,
			).Scan(&e.Id);
    
    if err != nil {
        return fmt.Errorf("Event.Store() error: %s", err)
    }

    return nil
}

func (e *Event) Update() error {
	return nil
}

func (e *Event) Exists() bool {
	db := GetConnection()
	err := db.QueryRow("SELECT TOP 1 FROM event WHERE link = ?;", e.Link).Scan()

	return err != sql.ErrNoRows
}

func (e *Event) IsDuplicate() bool {
	// db := GetConnection()
	return e.Exists()
}

func (e *Event) IsValid() bool {
	return e.Name != "" &&
		   e.Link != "" &&
		   e.Start.After(time.Now()) &&
		   e.Place.IsValid()
}