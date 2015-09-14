package model

import "strings"

type Tags []string

func (t *Tags) Encode() string {
    return "{" + strings.Join(*t, ",") + "}"
}

func (t *Tags) Decode(data string) error {
    *t = strings.Split(data[1:len(data)-1], ",")
    return nil
}