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

	// private
	err    error
}

/*func (q *EventQuery) Name(name string) *EventQuery {
	q.name = name
	return q
}

func (q *EventQuery) From(from time.Time) *EventQuery {
	q.from = from
	return q
}

func (q *EventQuery) To(to time.Time) *EventQuery {
	q.to = to
	return q
}

func (q *EventQuery) Position(lat float32, long float32) *EventQuery {
	q.lat = lat
	q.long = long
	return q
}

func (q *EventQuery) Tags(tags []string) *EventQuery {
	q.tags = tags
	return q
}

func (q *EventQuery) SortBy(sort SortType) *EventQuery {
	q.sort = sort
	return q
}

func (q *EventQuery) Limit(limit uint32) *EventQuery {
	q.limit = limit
	return q
}

func (q *EventQuery) Offset(offset uint32) *EventQuery {
	q.offset = offset
	return q
}*/

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
			var start, end, tags, gps string

			if err := rows.Scan(&e.Id, &e.Name, &e.Desc, &e.Link, &e.Image, &start, &end, &tags, &e.Place.Name, &gps); err != nil {
			    q.err = err
			    return
			}

			e.Start.Decode(start)
			e.End.Decode(end)
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
	sql.WriteString(" WHERE");
	needAnd := false

	if q.Name != "" {
		sql.WriteString(fmt.Sprintf(" name ~* '.*%s.*'", q.Name))
		needAnd = true
	}

	if q.From.IsValid() {
		if needAnd {
			sql.WriteString(" AND")
		}
		sql.WriteString(" starttime >= " + q.From.Encode())
	}

	if q.To.IsValid() {
		if needAnd {
			sql.WriteString(" AND")
		}
		sql.WriteString(" starttime <= " + q.To.Encode())
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY gps <-> point(%.4f,%.4f)", q.Lat, q.Long))
	
	if q.Limit > 0 {
		sql.WriteString(fmt.Sprintf(" LIMIT %d", q.Limit))
	}

	if q.Offset > 0 {
		sql.WriteString(fmt.Sprintf(" OFFSET %d", q.Offset))
	}

	sql.WriteString(";")

	return sql.String()
}
