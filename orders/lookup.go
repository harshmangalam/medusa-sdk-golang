package orders

type ShippingAddress struct {
	PostalCode string `json:"postal_code,omitempty" url:"postal_code,omitempty"`
}
type Lookup struct {
	DisplayId       string           `json:"display_id" url:"display_id"`
	Email           string           `json:"email" url:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}
