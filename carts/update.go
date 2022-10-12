package carts

type UpdateCart struct {
	RegionId         string         `json:"region_id"`
	country_code     string         `json:"country_code"`
	email            string         `json:"email"`
	sales_channel_id string         `json:"sales_channel_id"`
	billing_address  any            `json:"billing_address"`
	ipping_address   any            `json:"ipping_address"`
	gift_cards       any            `json:"gift_cards"`
	discounts        any            `json:"discounts"`
	customer_id      string         `json:"customer_id"`
	Context          map[string]any `json:"context"`
}
