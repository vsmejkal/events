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
	id, err := e.Exists()
	if err != nil {
		return fmt.Errorf("Event.Exists: %s", err)
	}

	if id > 0 {
		return e.update(id)
	} else {
		return e.insert()
	}
}

func (e *Event) Exists() (id int64, err error) {
	db := GetConnection()
	err = db.QueryRow("SELECT id FROM event WHERE link=$1 LIMIT 1;", e.Link).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil
	}
	return
}

func (e *Event) IsValid() bool {
	return e.Name != "" &&
		   e.Link != "" &&
		   e.Start.After(time.Now()) &&
		   e.Place.IsValid()
}

func (e *Event) insert() error {
	if err := e.Place.Store(); err != nil {
		return err
	}

	db := GetConnection()
	err := db.QueryRow("INSERT INTO event(name, description, link, image, starttime, endtime, tags, place) " +
	                   "VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id;",
				e.Name,
				e.Desc,
				e.Link,
				e.Image,
				e.Start.Encode(),
				e.End.Encode(),
				e.Tags.Encode(),
				e.Place.Id,
			).Scan(&e.Id);
    
    return err
}

func (e *Event) update(id int64) error {
	e.Id = id
	return nil
}