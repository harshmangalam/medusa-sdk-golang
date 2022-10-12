package carts

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type PaymentSession struct {
	ProviderId string `json:"provider_id"`
}

func CreatePaymentSession(cartId string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/payment-sessions", cartId)
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

func DeletePaymentSession(cartId, providerId string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/payment-sessions/%v", cartId, providerId)
	resp, err := request.NewRequest().SetMethod(http.MethodDelete).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func RefreshPaymentSession(cartId, providerId string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/payment-sessions/%v/refresh", cartId, providerId)
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

func NewPaymentSession() *PaymentSession {
	return new(PaymentSession)
}

func (p *PaymentSession) SetProviderId(providerId string) *PaymentSession {
	p.ProviderId = providerId
	return p
}
func (p *PaymentSession) Send(cartId string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/payment-session", cartId)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(p).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}