package orders

type ShippingAddress struct {
	PostalCode string `json:"postal_code,omitempty" url:"postal_code,omitempty"`
}
type Lookup struct {
	DisplayId       string           `json:"display_id" url:"display_id"`
	Email           string           `json:"email" url:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

func NewLookup() *Lookup {
	return new(Lookup)
}

func (l *Lookup) SetDisplayId(displayId string) *Lookup {
	l.DisplayId = displayId
	return l
}

func (l *Lookup) SetEmail(email string) *Lookup {
	l.Email = email
	return l
}

func (l *Lookup) SetShippingAddress(addr *ShippingAddress) *Lookup {
	l.ShippingAddress = addr
	return l
}
