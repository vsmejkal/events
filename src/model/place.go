package model

import (
	"fmt"
	"database/sql"
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
	id, err := p.Exists()
	if err != nil {
		return fmt.Errorf("Place.Exists: %s", err)
	}

	if id > 0 {
		return p.update(id)
	} else {
		return p.insert()
	}	
}

func (p Place) Exists() (id int64, err error) {
	db := GetConnection()
	err = db.QueryRow("SELECT id FROM place WHERE name=$1 LIMIT 1;", p.Name).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil
	}
	return
}

func (p Place) IsValid() bool {
    return p.Gps.IsValid()
}

func (p *Place) insert() error {
	db := GetConnection()
	err := db.QueryRow("INSERT INTO place(name, gps, street, city, zip, tags) " +
	                  "VALUES($1,$2,$3,$4,$5,$6) RETURNING id;", 
				p.Name,
				p.Gps.Encode(),
				p.Street,
				p.City,
				p.Zip,
				p.Tags.Encode(),
			).Scan(&p.Id);

    return err
}

func (p *Place) update(id int64) error {
	p.Id = id
	return nil
}