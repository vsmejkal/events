package model

type Place struct {
    Id uint64
	Title string
	Categories []string
    Latitude float64
    Longitude float64
	City string
    Street string
    Zip string
}