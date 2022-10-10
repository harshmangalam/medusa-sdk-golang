package carts

type CreateCartSchema struct {
	RegionId       string         `json:"region_id"`
	SalesChannelId string         `json:"sales_channel_id"`
	CountryCode    string         `json:"country_code"`
	Items          []*CartItem    `json:"items"`
	Context        map[string]any `json:"context"`
}
