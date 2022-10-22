package schema

import (
	"time"
)

type Customer struct {
	// The customer's email
	Email string `json:"email"`

	// The customer's ID
	Id string `json:"id"`

	// The customer's first name
	FirstName string `json:"first_name"`

	// The customer's last name
	LastName string `json:"last_name"`

	// The customer's billing address ID
	BillingAddressId string `json:"billing_address_id"`

	// An address.
	BillingAddress []*Address `json:"billing_address"`

	// Available if the relation shipping_addresses is expanded.
	ShippingAddress []*Address `json:"shipping_address"`

	// The customer's phone number
	Phone string `json:"phone"`

	// The customer's phone number
	HasAccount bool `json:"has_account"`

	// Available if the relation orders is expanded.
	Orders []*Order `json:"orders"`

	// The customer groups the customer belongs to. Available if the relation groups is expanded.
	Groups []*CustomerGroup `json:"groups"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata any `json:"metadata,omitempty"`
}
