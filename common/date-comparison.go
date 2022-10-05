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

// filter by dates greater than this date
func (d *DateComparison) SetGt(gt time.Time) *DateComparison {
	d.Gt = gt
	return d
}

// filter by dates greater than or equal to this date
func (d *DateComparison) SetGte(gte time.Time) *DateComparison {
	d.Gte = gte
	return d
}

// filter by dates less than or equal to this date
func (d *DateComparison) SetLte(lte time.Time) *DateComparison {
	d.Lte = lte
	return d
}
