package carts

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
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

func Complete(cart_id string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/complete", cart_id)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
