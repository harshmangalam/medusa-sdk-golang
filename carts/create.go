package carts

type CartItem struct {
	VariantId string `json:"variant_id"`
	Quantity  int    `json:"quantity"`
}
type CreateCartSchema struct {
	RegionId       string         `json:"region_id"`
	SalesChannelId string         `json:"sales_channel_id"`
	CountryCode    string         `json:"country_code"`
	Items          []*CartItem    `json:"items"`
	Context        map[string]any `json:"context"`
}
