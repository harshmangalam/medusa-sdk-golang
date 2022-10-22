package orderedits

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

type DeclineData struct {
	// Gift Cards are redeemable and represent a value that can be used towards the payment of an Order.
	OrderEdit *schema.OrderEdit `json:"order_edit"`
}

type DeclineResponse struct {
	// Success response
	Data *DeclineData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type DeclineBody struct {
	DeclinedReason string `json:"declined_reason,omitempty"`
}

func NewDeclineBody() *DeclineBody {
	return new(DeclineBody)
}

func (d *DeclineBody) SetDeclineReason(reason string) *DeclineBody {
	d.DeclinedReason = reason
	return d
}

func (d *DeclineBody) Decline(id string, config *medusa.Config) (*DeclineResponse, error) {
	path := fmt.Sprintf("/store/order-edits/%v/decline", id)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(d).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(DeclineResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(DeclineData)
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
