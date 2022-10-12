package carts

import (
	"fmt"
	"net/http"
	"time"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type Discount struct {
	Code string `json:"code"`
}

type GiftCard struct {
	Code string `json:"code"`
}

type Address struct {
	Id          string         `json:"id"`
	CustomerId  string         `json:"customer_id"`
	Customer    []any          `json:"customer"`
	FirstName   []any          `json:"first_name"`
	LastName    []any          `json:"last_name"`
	Address1    []any          `json:"address_1"`
	Address2    []any          `json:"address_2"`
	City        []any          `json:"city"`
	CountryCode []any          `json:"country_code"`
	Country     any            `json:"country"`
	Province    string         `json:"province"`
	PostalCode  string         `json:"postal_code"`
	Phone       string         `json:"phone"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
	DeletedAt   time.Time      `json:"deleted_at,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}
type UpdateCart struct {
	RegionId        string         `json:"region_id"`
	CountryCode     string         `json:"country_code"`
	Email           string         `json:"email"`
	SalesChannelId  string         `json:"sales_channel_id"`
	BillingAddress  any            `json:"billing_address"`
	ShippingAddress any            `json:"ipping_address"`
	GiftCards       []*GiftCard    `json:"gift_cards"`
	Discounts       []*Discount    `json:"discounts"`
	CustomerId      string         `json:"customer_id"`
	Context         map[string]any `json:"context"`
}

func NewUpdateCart() *UpdateCart {
	u := new(UpdateCart)
	return u
}

func (u *UpdateCart) SetRegionId(regionId string) *UpdateCart {
	u.RegionId = regionId
	return u
}

func (u *UpdateCart) SetCountryCode(countryCode string) *UpdateCart {
	u.CountryCode = countryCode
	return u
}

func (u *UpdateCart) SetEmail(email string) *UpdateCart {
	u.Email = email
	return u
}
func (u *UpdateCart) SetSalesChannelId(salesChannelId string) *UpdateCart {
	u.SalesChannelId = salesChannelId
	return u
}

func (u *UpdateCart) SetBillingAddress(address any) *UpdateCart {
	u.BillingAddress = address
	return u
}

func (u *UpdateCart) SetShippingAddress(address any) *UpdateCart {
	u.ShippingAddress = address
	return u
}

func (u *UpdateCart) SetGiftCards(giftCards []*GiftCard) *UpdateCart {
	u.GiftCards = giftCards
	return u
}

func (u *UpdateCart) SetDiscounts(discounts []*Discount) *UpdateCart {
	u.Discounts = discounts
	return u
}

func (u *UpdateCart) SetCustomerId(customerId string) *UpdateCart {
	u.CustomerId = customerId
	return u
}

func (u *UpdateCart) SetContext(context map[string]any) *UpdateCart {
	u.Context = context
	return u
}

func (u *UpdateCart) Update(cartId string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/carts/%v", cartId)
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
