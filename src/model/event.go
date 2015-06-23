package model

import "time"

type Event struct {
    Id uint32
	Title string
	Desc string
	Link string
	Image string
	Categories []string
	Start time.Time
	End time.Time
	IsDateOnly bool
	Place Place
}