package common

import "time"

type DateComparisonOperator struct {
	Lt  time.Time `json:"lt,omitempty"`
	Gt  time.Time `json:"gt,omitempty"`
	Lte time.Time `json:"lte,omitempty"`
	Gte time.Time `json:"gte,omitempty"`
}
