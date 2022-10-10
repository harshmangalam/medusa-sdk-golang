package carts

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type CartItem struct {
	VariantId string `json:"variant_id"`
	Quantity  int    `json:"quantity"`
}
type CreateCart struct {
	RegionId       string         `json:"region_id"`
	SalesChannelId string         `json:"sales_channel_id"`
	CountryCode    string         `json:"country_code"`
	Items          []*CartItem    `json:"items"`
	Context        map[string]any `json:"context"`
}

func NewCreateCart() *CreateCart {
	return new(CreateCart)
}
func (c *CreateCart) Create(config *medusa.Config) ([]byte, error) {
	const path = `/store/carts`
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetData(c).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
