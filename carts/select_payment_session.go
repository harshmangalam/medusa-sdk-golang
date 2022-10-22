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

type SelectPaymentSessionData struct {
	Cart *schema.Cart `json:"cart"`
}

type SelectPaymentSessionResponse struct {
	// Success response
	Data *SelectPaymentSessionData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type SelectPaymentSession struct {
	ProviderId string `json:"provider_id"`
}

func NewSelectPaymentSession() *SelectPaymentSession {
	return new(SelectPaymentSession)
}

func (p *SelectPaymentSession) SetProviderId(providerId string) *SelectPaymentSession {
	p.ProviderId = providerId
	return p
}

// Selects a Payment Session as the session intended to be used towards the completion of the Cart.
func (p *SelectPaymentSession) Select(cartId string, config *medusa.Config) (*SelectPaymentSessionResponse, error) {
	path := fmt.Sprintf("/store/carts/%v/payment-session", cartId)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(p).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(SelectPaymentSessionResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(SelectPaymentSessionData)
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
