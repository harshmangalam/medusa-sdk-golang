package swaps

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

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

func (c *CreateSwap) Create(config *medusa.Config) (*Swap, error) {
	path := "/store/swaps"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(c).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(ResponseBody)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}

	return respBody.Swap, nil
}
