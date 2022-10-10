package carts

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ShippingMethodSchema struct {
	OptionId string `json:"option_id"`
	Data     string `json:"data"`
}

func NewShippingMethodSchema() *ShippingMethodSchema {
	return new(ShippingMethodSchema)
}

func (s *ShippingMethodSchema) SetOptionId(optionId string) *ShippingMethodSchema {
	s.OptionId = optionId
	return s
}

func (s *ShippingMethodSchema) SetData(data string) *ShippingMethodSchema {
	s.Data = data
	return s
}

func (s *ShippingMethodSchema) Save(cart_id string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/shipping-methods", cart_id)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(s).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
