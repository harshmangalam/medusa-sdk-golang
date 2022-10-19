package swaps

type CreateSwap struct {
	OrderId              string `json:"order_id"`
	ReturnItems          []any  `json:"return_items"`
	AdditionalItems      []any  `json:"additional_items"`
	ReturnShippingOption string `json:"return_shipping_option"`
}

func NewCreateSwap() *CreateSwap {
	return new(CreateSwap)
}

func (c *CreateSwap) SetOrderId(orderId string) *CreateSwap {
	c.OrderId = orderId
	return c
}

func (c *CreateSwap) SetReturnItems(items []any) *CreateSwap {
	c.ReturnItems = items
	return c
}
func (c *CreateSwap) SetAdditionalItems(items []any) *CreateSwap {
	c.AdditionalItems = items
	return c
}
func (c *CreateSwap) SetReturnShippingOption(shippingOpions string) *CreateSwap {
	c.ReturnShippingOption = shippingOpions
	return c
}
