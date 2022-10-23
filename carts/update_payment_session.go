package carts

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

type UpdatePaymentSessionData struct {
	Cart *schema.Cart `json:"cart"`
}

type UpdatePaymentSessionResponse struct {
	// Success response
	Data *UpdatePaymentSessionData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type UpdatePaymentSession struct {
	Data map[string]any `json:"data"`
}

func NewUpdatePaymentSession() *UpdatePaymentSession {
	return new(UpdatePaymentSession)
}

func (u *UpdatePaymentSession) SetData(data map[string]any) *UpdatePaymentSession {
	u.Data = data
	return u
}

func (u *UpdatePaymentSession) Update(cartId, providerId string, config *medusa.Config) (*UpdatePaymentSessionResponse, error) {
	path := fmt.Sprintf("/store/carts/%v/payment-sessions/%v", cartId, providerId)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(u).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(UpdatePaymentSessionResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(UpdatePaymentSessionData)
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
