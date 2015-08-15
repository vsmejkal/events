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
	Tags     []int
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
	_, err := db.Exec("INSERT INTO event(name, description, link, starttime, endtime) VALUES($1,$2,$3,$4);", "XXX", "YYY", "2012-12-09", "2012-12-09")
    if (err != nil) {
        fmt.Println(err)
        return
    }
}

func (e *Event) IsDuplicate() bool {
	db := GetConnection()
	return false
}

func (e *Event) IsValid() bool {
	return e.Name != "" &&
		e.Link != "" &&
		e.Start.After(time.Now()) &&
		e.End.After(time.Now()) &&
		e.Place.IsValid()
}
