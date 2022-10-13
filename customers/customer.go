package customers

import (
	"time"

	"github.com/harshmngalam/medusa-sdk-golang/address"
)

type Customer struct {
	Email            string             `json:"email"`
	Id               string             `json:"id"`
	FirstName        string             `json:"first_name"`
	LastName         string             `json:"last_name"`
	BillingAddressId string             `json:"billing_address_id"`
	BillingAddress   []*address.Address `json:"billing_address"`
	ShippingAddress  []*address.Address `json:"shipping_address"`
	Phone            string             `json:"phone"`
	HasAccount       bool               `json:"has_account"`
	Orders           []any              `json:"orders"`
	Groups           []*CustomerGroup   `json:"groups"`
	CreatedAt        time.Time          `json:"created_at,omitempty"`
	UpdatedAt        time.Time          `json:"updated_at,omitempty"`
	DeletedAt        time.Time          `json:"deleted_at,omitempty"`
	Metadata         any                `json:"metadata,omitempty"`
}
