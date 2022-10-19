package swaps

type CreateSwap struct {
	OrderId              string `json:"order_id"`
	ReturnItems          []any  `json:"return_items"`
	AdditionalItems      []any  `json:"additional_items"`
	ReturnShippingOption string `json:"return_shipping_option"`
}
