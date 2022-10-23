package carts

import (
	"encoding/json"
	"fmt"
	"net/http"

	medusa "github.com/harshmangalam/medusa-sdk-golang"
	"github.com/harshmangalam/medusa-sdk-golang/request"
	"github.com/harshmangalam/medusa-sdk-golang/response"
	"github.com/harshmangalam/medusa-sdk-golang/utils"
)

type CompleteDataEnum string

const (
	CompleteOrder CompleteDataEnum = "order"
	CompleteCart  CompleteDataEnum = "cart"
	CompleteSwap  CompleteDataEnum = "swap"
)

type CompleteData struct {
	// The type of the data property.
	Type CompleteDataEnum `json:"type"`

	// Data will be one of order, cart and swap
	Data any `json:"data"`
}

type CompleteResponse struct {
	// Success response
	Data *CompleteData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

// Completes a cart. The following steps will be performed. Payment authorization is attempted and if more work is required, we simply return the cart for further updates. If payment is authorized and order is not yet created, we make sure to do so. The completion of a cart can be performed idempotently with a provided header Idempotency-Key. If not provided, we will generate one for the request.

func Complete(cartId string, config *medusa.Config) (*CompleteResponse, error) {
	path := fmt.Sprintf("/store/carts/%v/complete", cartId)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(CompleteResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(CompleteData)
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
