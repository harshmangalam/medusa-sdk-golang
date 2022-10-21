package carts

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ShippingMethod struct {
	OptionId string `json:"option_id"`
	Data     string `json:"data"`
}

func NewShippingMethod() *ShippingMethod {
	return new(ShippingMethod)
}

func (s *ShippingMethod) SetOptionId(optionId string) *ShippingMethod {
	s.OptionId = optionId
	return s
}

func (s *ShippingMethod) SetData(data string) *ShippingMethod {
	s.Data = data
	return s
}

func (s *ShippingMethod) Add(cartId string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/shipping-methods", cartId)
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
