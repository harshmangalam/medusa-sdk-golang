package common

import "time"

type DateComparison struct {
	Lt  time.Time `json:"lt,omitempty"`
	Gt  time.Time `json:"gt,omitempty"`
	Lte time.Time `json:"lte,omitempty"`
	Gte time.Time `json:"gte,omitempty"`
}

func NewDateComparison() *DateComparison {
	return new(DateComparison)
}
