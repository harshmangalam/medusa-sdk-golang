package carts

type UpdateCart struct {
	RegionId       string         `json:"region_id"`
	CountryCode    string         `json:"country_code"`
	Email          string         `json:"email"`
	SalesChannelId string         `json:"sales_channel_id"`
	BillingAddress any            `json:"billing_address"`
	SippingAddress any            `json:"ipping_address"`
	GiftCards      any            `json:"gift_cards"`
	Discounts      any            `json:"discounts"`
	CustomerId     string         `json:"customer_id"`
	Context        map[string]any `json:"context"`
}
