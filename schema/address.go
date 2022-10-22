package schema

import "time"

type Address struct {
	Id          string         `json:"id"`
	CustomerId  string         `json:"customer_id"`
	Customer    []any          `json:"customer"`
	FirstName   []any          `json:"first_name"`
	LastName    []any          `json:"last_name"`
	Address1    []any          `json:"address_1"`
	Address2    []any          `json:"address_2"`
	City        []any          `json:"city"`
	CountryCode []any          `json:"country_code"`
	Country     any            `json:"country"`
	Province    string         `json:"province"`
	PostalCode  string         `json:"postal_code"`
	Phone       string         `json:"phone"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
	DeletedAt   time.Time      `json:"deleted_at,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}
