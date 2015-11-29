package model

import "time"

type Source struct {
	Id int64
	Url string
	Place Place
	Visited time.Time
}