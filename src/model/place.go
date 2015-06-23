package model

type Place struct {
    Id uint32
	Title string
	Categories []string
    Latitude float32
    Longitude float32
	City string
    Street string
    Zip int
}