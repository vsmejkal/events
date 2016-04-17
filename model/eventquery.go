package model

import (
	"bytes"
	"fmt"
)

type SortType int
const (
	DISTANCE SortType = iota
	IMPORTANCE
)

type EventQuery struct {
	Name   string
	From   Datetime
	To     Datetime
	Lat    float32
	Long   float32
	Tags   []string
	Sort   SortType
	Limit  uint32
	Offset uint32

	err    error
}


func (q *EventQuery) Search() <-chan Event {
	out := make(chan Event)
	
	go func() {
		defer close(out)

		sql := q.composeSQL()
		db := GetConnection()
		rows, err := db.Query(sql)
		if err != nil {
			q.err = err
			return
		}

		for rows.Next() {
			var e Event
			var tags, gps string
			if err := rows.Scan(&e.Id, &e.Name, &e.Desc, &e.Link, &e.Image, &e.Start.Time, &e.End.Time, &tags, &e.Place.Name, &gps); err != nil {
			    q.err = err
			    return
			}

			e.Tags.Decode(tags)
			e.Place.Gps.Decode(gps)
			out <- e
		}

        q.err = rows.Err()
    }()

	return out
}

func (q *EventQuery) Error() error {
	return q.err
}

func (q *EventQuery) composeSQL() string {
	var sql bytes.Buffer

	sql.WriteString(" SELECT e.id, e.name, e.description, e.link, e.image, e.starttime, e.endtime, e.tags, p.name, p.gps")
	sql.WriteString(" FROM event e INNER JOIN place p ON e.place = p.id")
	sql.WriteString(" WHERE 1=1");

	if q.Name != "" {
		sql.WriteString(fmt.Sprintf(" AND name ~* '.*%s.*'", q.Name))
	}

	if !q.From.IsZero() {
		sql.WriteString(" AND starttime >= '" + q.From.Value() + "'")
	}

	if !q.To.IsZero() {
		sql.WriteString(" AND starttime <= '" + q.To.Value() + "'")
	}

	sql.WriteString(" ORDER BY starttime")
	// sql.WriteString(fmt.Sprintf(" ORDER BY gps <-> point(%.4f,%.4f)", q.Lat, q.Long))
	
	if q.Limit > 0 {
		sql.WriteString(fmt.Sprintf(" LIMIT %d", q.Limit))
	}

	if q.Offset > 0 {
		sql.WriteString(fmt.Sprintf(" OFFSET %d", q.Offset))
	}

	sql.WriteString(";")

	// fmt.Println(sql.String())

	return sql.String()
}
