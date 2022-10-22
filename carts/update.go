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

type UpdateCartData struct {
	Cart *schema.Cart `json:"cart"`
}

type UpdateCartResponse struct {
	// Success response
	Data *UpdateCartData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type Discount struct {
	Code string `json:"code"`
}

type GiftCard struct {
	Code string `json:"code"`
}

type UpdateCart struct {
	// The id of the Region to create the Cart in.
	RegionId string `json:"region_id"`

	// The 2 character ISO country code to create the Cart in.
	CountryCode string `json:"country_code"`

	// An email to be used on the Cart.
	Email string `json:"email"`

	// The ID of the Sales channel to update the Cart with.
	SalesChannelId string `json:"sales_channel_id"`

	// The Address to be used for billing purposes.
	BillingAddress any `json:"billing_address"`

	// The Address to be used for shipping.
	ShippingAddress any `json:"ipping_address"`

	// An array of Gift Card codes to add to the Cart.
	GiftCards []*GiftCard `json:"gift_cards"`

	// An array of Discount codes to add to the Cart.
	Discounts []*Discount `json:"discounts"`

	// The ID of the Customer to associate the Cart with.
	CustomerId string `json:"customer_id"`

	// An optional object to provide context to the Cart.
	Context map[string]any `json:"context"`
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

// Updates a Cart.
func (u *UpdateCart) Update(cartId string, config *medusa.Config) (*UpdateCartResponse, error) {
	path := fmt.Sprintf("/store/carts/%v", cartId)
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(u).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(UpdateCartResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(UpdateCartData)
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
