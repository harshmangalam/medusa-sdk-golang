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

type UpdatePaymentSession struct {
	Data map[string]any `json:"data"`
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

func NewUpdatePaymentSession() *UpdatePaymentSession {
	return new(UpdatePaymentSession)
}

func (u *UpdatePaymentSession) SetData(data map[string]any) *UpdatePaymentSession {
	u.Data = data
	return u
}

func (u *UpdatePaymentSession) Apply(cartId, providerId string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v/payment-sessions/%v", cartId, providerId)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(u).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
