package model

import "time"

type Datetime struct {
	time.Time
}

func (d *Datetime) Encode() string {
	return d.Format(time.RFC3339)
}

func (d *Datetime) Decode(data string) error {
	t, err := time.Parse(time.RFC3339, data)
	*d = Datetime{t}
	return err
}

// func (d Datetime) IsValid() bool {
//	 return d != Datetime{time.Time{}}
// }

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
