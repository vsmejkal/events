package model

import (
	"time"
	"errors"
	"database/sql/driver"
)

type Datetime struct {
	time.Time
}

func (d Datetime) Value() (driver.Value, error) {
	return []byte(d.Format(time.RFC3339)), nil
}

// Implements sql.Scanner interface
func (d *Datetime) Scan(src interface{}) error {
	var source string

	if (src == nil) {
		*d = Datetime{}
		return nil
	}

	switch src.(type) {
	case string, []byte:
		source = src.(string)
	default:
		return errors.New("Incompatible type for Datetime")
	}

	t, err := time.Parse(time.RFC3339, source)
	*d = Datetime{t}
	return err
}

 func (d Datetime) Valid() bool {
	 return d.Unix() > 0
 }

func (d Datetime) HumanDate() string {
	today := time.Now()
	tomorrow := today.Add(time.Duration(24) * time.Hour)

	if today.Day() == d.Day() && today.Month() == d.Month() && today.Year() == d.Year() {
		return "dnes"
	}

	if tomorrow.Day() == d.Day() && tomorrow.Month() == d.Month() && tomorrow.Year() == d.Year() {
		return "zítra"
	}

	return d.Format("02/01")
}
