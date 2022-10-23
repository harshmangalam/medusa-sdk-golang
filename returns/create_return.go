package returns

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmangalam/medusa-sdk-golang"
	"github.com/harshmangalam/medusa-sdk-golang/request"
	"github.com/harshmangalam/medusa-sdk-golang/response"
	"github.com/harshmangalam/medusa-sdk-golang/schema"
	"github.com/harshmangalam/medusa-sdk-golang/utils"
)

type CreateReturnData struct {
	Return []*schema.Return `json:"return"`
}

type CreateReturnResponse struct {
	// Success response
	Data *CreateReturnData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type CreateReturn struct {
	// The ID of the Order to create the Return from.

	OrderId string `json:"order_id"`

	// The items to include in the Return.
	Items []any `json:"items"`

	// If the Return is to be handled by the store operator the Customer can choose a Return Shipping Method. Alternatvely the Customer can handle the Return themselves.
	ReturnShipping any `json:"return_shipping"`
}

func NewCreateRetun() *CreateReturn {
	return new(CreateReturn)
}

func (c *CreateReturn) SetOrderId(orderId string) *CreateReturn {
	c.OrderId = orderId
	return c
}
func (c *CreateReturn) SetItems(items []any) *CreateReturn {
	c.Items = items
	return c

}

func (c *CreateReturn) SetReturnShipping(shipping any) *CreateReturn {
	c.ReturnShipping = shipping
	return c
}

func (c *CreateReturn) Create(config *medusa.Config) (*CreateReturnResponse, error) {
	path := "/store/returns"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(c).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(CreateReturnResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(CreateReturnData)
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
