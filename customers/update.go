package customers

import (
	"github.com/harshmngalam/medusa-sdk-golang/address"
)

type UpdateCustomer struct {
	Email          string             `json:"email"`
	FirstName      string             `json:"first_name"`
	LastName       string             `json:"last_name"`
	BillingAddress []*address.Address `json:"billing_address"`
	Phone          string             `json:"phone"`
	Password       string             `json:"password"`
	Metadata       any                `json:"metadata,omitempty"`
}
