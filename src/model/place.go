package model

import (
	"fmt"
)

type Place struct {
    Id      int64
	Name    string
    Gps     Gps
    Street  string
    City    string
    Zip     string
    Tags    Tags
}

func (p *Place) Store() error {
	db := GetConnection()
	err := db.QueryRow("INSERT INTO place(name, gps, street, city, zip, tags) VALUES($1,$2,$3,$4,$5,$6) RETURNING id;", 
				p.Name,
				p.Gps.Encode(),
				p.Street,
				p.City,
				p.Zip,
				p.Tags.Encode(),	
			).Scan(&p.Id);

	if err != nil {
		return fmt.Errorf("Place.Store() error: %s", err)
	}

    return nil
}

func (p *Place) IsValid() bool {
    return p.Gps.IsValid()
}