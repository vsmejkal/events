package model

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
	"database/sql/driver"
)

type GPS struct {
	Lat  float64
	Long float64
}

func (gps GPS) Value() (driver.Value, error) {
	return []byte(fmt.Sprintf("%.8f,%.8f", gps.Lat, gps.Long)), nil
}

func (gps *GPS) Decode(src interface{}) error {
	var data string

	if (src == nil) {
		*gps = GPS{}
		return nil
	}

	switch src.(type) {
	case string, []byte:
		data = src.(string)
	default:
		return errors.New("Incompatible type for GPS")
	}

	fields := strings.Split(data, ",")
	if len(fields) != 2 {
		return errors.New("GPS: Decode error (need 2 items, got " + string(len(fields)) + ")")
	}

	var err error
	if gps.Lat, err = strconv.ParseFloat(fields[0], 32); err != nil {
		return err
	}
	if gps.Long, err = strconv.ParseFloat(fields[1], 32); err != nil {
		return err
	}

	return nil
}

func (gps GPS) Valid() bool {
	return gps.Lat != 0 && gps.Long != 0
}