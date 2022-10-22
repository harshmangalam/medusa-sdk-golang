package schema

import "time"

type Address struct {
	// ID of the address
	Id string `json:"id"`

	// ID of the customer this address belongs to
	CustomerId string `json:"customer_id"`

	// Available if the relation customer is expanded.
	Customer []any `json:"customer"`
	// Company name
	Company string `json:"company,omitempty"`

	// First name
	FirstName []any `json:"first_name"`

	// Last name
	LastName []any `json:"last_name"`

	// Address line 1
	Address1 []any `json:"address_1"`

	// Address line 2
	Address2 []any `json:"address_2"`

	// City
	City []any `json:"city"`

	// The 2 character ISO code of the country in lower case
	CountryCode []any `json:"country_code"`

	// A country object. Available if the relation country is expanded.
	Country any `json:"country"`

	// Province
	Province string `json:"province"`

	// Postal Code
	PostalCode string `json:"postal_code"`

	// Phone Number
	Phone string `json:"phone"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata map[string]any `json:"metadata,omitempty"`
}
