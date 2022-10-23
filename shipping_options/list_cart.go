package shippingoptions

import (
	"encoding/json"
	"fmt"
	"net/http"

	medusa "github.com/harshmangalam/medusa-sdk-golang"
	"github.com/harshmangalam/medusa-sdk-golang/request"
	"github.com/harshmangalam/medusa-sdk-golang/response"
	"github.com/harshmangalam/medusa-sdk-golang/schema"
	"github.com/harshmangalam/medusa-sdk-golang/utils"
)

type ListCartData struct {
	ShippingOptions []*schema.ShippingOption `json:"shipping_options"`
}

type ListCartResponse struct {
	// Success response
	Data *ListCartData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

// Retrieves a list of Shipping Options available to a cart.
func ListCartOptions(cartId string, config *medusa.Config) (*ListCartResponse, error) {
	path := fmt.Sprintf("store/shipping-options/%v", cartId)

	resp, err := request.
		NewRequest().
		SetMethod(http.MethodGet).
		SetPath(path).
		Send(config)

	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(ListCartResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ListCartData)
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
