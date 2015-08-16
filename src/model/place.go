package model

import (
	"fmt"
)

type Place struct {
    Id int64
	Name string
    Lat float64
    Long float64
    Street string
    City string
    Zip string
	Tags []string
}

func (p *Place) Store() error {
	db := GetConnection()

	err := db.QueryRow("INSERT INTO place(name, gps, street, city, zip, tags) VALUES($1,$2,$3,$4,$5,$6) RETURNING id;", 
				p.Name,
				fmt.Sprintf("%.8f,%.8f", p.Lat, p.Long),
				p.Street,
				p.City,
				p.Zip,
				SerializeStringArray(p.Tags),	
			).Scan(&p.Id);

	if err != nil {
		return fmt.Errorf("Place.Store() error: %s", err)
	}

    return nil
}

func (p *Place) IsValid() bool {
    return p.Lat != 0 && p.Long != 0
}