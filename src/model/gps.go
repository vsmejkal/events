package model

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
)

type Gps struct {
	Lat  float64
	Long float64
}

func (gps *Gps) Encode() string {
	return fmt.Sprintf("%.8f,%.8f", gps.Lat, gps.Long)
}

func (gps *Gps) Decode(data string) error {
	fields := strings.Split(data, ",")
	if len(fields) != 2 {
		return errors.New("Gps.Decode error: need 2 items, got " + string(len(fields)))
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

func (gps *Gps) IsValid() bool {
	return gps.Lat != 0 && gps.Long != 0
}