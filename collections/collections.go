package collections

import (
	"time"
)

type Collection struct {
	// The title that the Product Collection is identified by.
	Title string `json:"title"`

	// The product collection's ID
	Id string `json:"id,omitempty"`

	// A unique string that identifies the Product Collection - can for example be used in slug structures.
	Handle string `json:"handle,omitempty"`

	// The Products contained in the Product Collection. Available if the relation products is expanded.
	Products []any `json:"products,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata any `json:"metadata,omitempty"`
}
