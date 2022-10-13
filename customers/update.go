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

func NewUpdateCustomer() *UpdateCustomer {

	return new(UpdateCustomer)
}
func (u *UpdateCustomer) SetEmail(email string) *UpdateCustomer {
	u.Email = email
	return u
}
func (u *UpdateCustomer) SetFirstName(firstName string) *UpdateCustomer {
	u.FirstName = firstName
	return u
}
