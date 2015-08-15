package model

import (
	"fmt"
	"strings"
)

type Place struct {
    Id uint64
	Name string
    Lat float64
    Long float64
    Street string
    City string
    Zip string
	Tags []int
}

func (p *Place) Store() error {
	db := GetConnection()

	db.Exec("INSERT INTO place(name, gps, street, city, zip, tags) VALUES($1, $2, $3, $4, $5, $6);", 
		p.Name,
		fmt.Sprintf("%.8f,%.8f", p.Lat, p.Long),
		p.Street,
		p.City,
		p.Zip,
		SerializeInts(p.Tags)
	)
}

func (p *Place) IsValid() bool {
    return p.Lat != 0 && p.Long != 0
}