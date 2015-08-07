package model

type Place struct {
    Id uint64
	Name string
    Lat float64
    Long float64
    Street string
    City string
    Zip string
	Tags []string
}

func (p *Place) IsValid() bool {
    return p.Lat != 0 && p.Long != 0
}