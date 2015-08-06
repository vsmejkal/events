package model

import "time"

type Source struct {
    Id uint64
    Url string
    Place Place
    Visited time.Time
}