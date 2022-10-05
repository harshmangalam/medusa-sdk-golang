package common

import "time"

type DateComparison struct {
	Lt  time.Time `json:"lt,omitempty"`
	Gt  time.Time `json:"gt,omitempty"`
	Lte time.Time `json:"lte,omitempty"`
	Gte time.Time `json:"gte,omitempty"`
}

// create new date comparison
func NewDateComparison() *DateComparison {
	return new(DateComparison)
}

// filter by dates less than this date
func (d *DateComparison) SetLt(lt time.Time) *DateComparison {
	d.Lt = lt
	return d
}
