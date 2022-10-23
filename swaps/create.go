package swaps

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmangalam/medusa-sdk-golang"
	"github.com/harshmangalam/medusa-sdk-golang/request"
	"github.com/harshmangalam/medusa-sdk-golang/response"
	"github.com/harshmangalam/medusa-sdk-golang/schema"
	"github.com/harshmangalam/medusa-sdk-golang/utils"
)

type CreateSwapData struct {
	Swap *schema.Swap `json:"swap"`
}

type CreateSwapResponse struct {
	// Success response
	Data *CreateSwapData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type CreateSwap struct {
	// The ID of the Order to create the Swap for.
	OrderId string `json:"order_id"`

	// The items to include in the Return.
	ReturnItems []any `json:"return_items"`

	// The items to exchange the returned items to.
	AdditionalItems []any `json:"additional_items"`

	// The ID of the Shipping Option to create the Shipping Method from.
	ReturnShippingOption string `json:"return_shipping_option,omitempty"`
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

// Creates a Swap on an Order by providing some items to return along with some items to send back

func (c *CreateSwap) Create(config *medusa.Config) (*CreateSwapResponse, error) {
	path := "/store/swaps"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(c).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(CreateSwapResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(CreateSwapData)
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
