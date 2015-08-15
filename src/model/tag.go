package model

import (
	"bytes"
	"strings"
	"strconv"
)

type Tag struct {
	Id uint64
	Name string
	Label string
}

func LoadTags() ([]Tag, error) {

}

func SerializeInts(arr []int) string {
	var buf bytes.Buffer
	for _, n := range arr  {
	    buf.WriteString(strconv.Itoa(n))
	    buf.WriteString(",")
	}

	// Remove trailing comma
	buf.Truncate(buf.Len() - 1)

	return buf.String()
}

func DeserializeInts(data string) []int {
	strs := strings.Split(data, ",")
	ints := make([]int, len(strs), len(strs))

	for i, n := range strs {
		ints[i], _ = strconv.Atoi(n)
	}

	return ints
}