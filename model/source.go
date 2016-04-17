package model

import (
	"time"
	"database/sql"
)

type Source struct {
	Id int64
	Name string
	Url string
	Place Place
	Visited Datetime
}