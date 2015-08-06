package model

import (
	"bytes"
	"errors"
	"time"
)

type SortType int

const (
	DISTANCE SortType = iota
	IMPORTANCE
)

type EventQuery struct {
	Range float32
	Start time.Time
	End   time.Time
	Sort  SortType
	Name  string
	Tags  []string
}

//
func (eq *EventQuery) Search() (events []Event, err error) {
	if eq.Range <= 0 {
		err = errors.New("Range must be greater than zero")
		return
	}

	if !eq.Start.Before(eq.End) {
		err = errors.New("Start time must be before end time")
		return
	}

	// sql := eq.composeSQL();
	// execute SQL
	// parse results

	return
}

func (eq *EventQuery) composeSQL() string {
	var sql bytes.Buffer

	// sql.WriteString()

	return sql.String()
}
