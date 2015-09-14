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

func (d *Datetime) IsValid() bool {
    return *d != Datetime{time.Time{}}
}