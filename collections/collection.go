package collections

import (
	"time"
)

type Collection struct {
	Title     string    `json:"title"`
	Id        string    `json:"id,omitempty"`
	Handle    string    `json:"handle,omitempty"`
	Products  []any     `json:"products,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Metadata  any       `json:"metadata,omitempty"`
}
