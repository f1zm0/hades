package hashing

import (
	"strings"
)

type DJB2 struct{}

func NewDJB2() *DJB2 {
	return &DJB2{}
}

func (d *DJB2) HashString(s string) int64 {
	var hash int64 = 5381
	for _, c := range strings.ToLower(s) {
		hash = ((hash << 5) + hash) + int64(c)
	}
	return hash
}
