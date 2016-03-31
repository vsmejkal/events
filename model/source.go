package model

import (
	"time"
	"database/sql"
)

type Source struct {
	Id int64
	Url string
	Place Place
	Visited time.Time
}