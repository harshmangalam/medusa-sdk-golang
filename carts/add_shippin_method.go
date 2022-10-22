package carts

import (
	"encoding/json"
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ShippingMethodData struct {
	Cart *schema.Cart `json:"cart"`
}

type ShippingMethodResponse struct {
	// Success response
	Data *ShippingMethodData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

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

// Adds a Shipping Method to the Cart.
func (s *ShippingMethod) Add(cartId string, config *medusa.Config) (*ShippingMethodResponse, error) {
	path := fmt.Sprintf("/store/carts/%v/shipping-methods", cartId)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(s).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(ShippingMethodResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ShippingMethodData)
		if err := json.Unmarshal(body, respData); err != nil {
			return nil, err
		}
		respBody.Data = respData

	case http.StatusUnauthorized:
		respErr := utils.UnauthorizeError()
		respBody.Error = respErr

	case http.StatusBadRequest:
		respErrors, err := utils.ParseErrors(body)
		if err != nil {
			return nil, err
		}
		if len(respErrors.Errors) == 0 {
			respError, err := utils.ParseError(body)
			if err != nil {
				return nil, err
			}
			respBody.Error = respError
		} else {
			respBody.Errors = respErrors
		}

	default:
		respErr, err := utils.ParseError(body)
		if err != nil {
			return nil, err
		}
		respBody.Error = respErr
	}

	return respBody, nil
}
