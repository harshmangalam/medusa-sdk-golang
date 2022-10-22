package carts

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type CartContext map[string]any

type CartItem struct {
	VariantId string `json:"variant_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}
type CreateCart struct {
	RegionId       string      `json:"region_id,omitempty"`
	SalesChannelId string      `json:"sales_channel_id,omitempty"`
	CountryCode    string      `json:"country_code,omitempty"`
	Items          []*CartItem `json:"items,omitempty"`
	Context        CartContext `json:"context,omitempty"`
}

type CreateData struct {
	Cart *schema.Cart `json:"cart"`
}

type CreateResponse struct {
	// Success response
	Data *CreateData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

func NewCreateCart() *CreateCart {
	return new(CreateCart)
}

func (c *CreateCart) SetRegionId(regionId string) *CreateCart {
	c.RegionId = regionId
	return c
}
func (c *CreateCart) SetSalesChannelId(salesChannelId string) *CreateCart {
	c.SalesChannelId = salesChannelId
	return c
}
func (c *CreateCart) SetCountryCode(countryCode string) *CreateCart {
	c.CountryCode = countryCode
	return c
}
func (c *CreateCart) SetItems(items []*CartItem) *CreateCart {
	c.Items = items
	return c
}
func (c *CreateCart) SetContext(context CartContext) *CreateCart {
	c.Context = context
	return c
}

// Creates a Cart within the given region and with the initial items. If no region_id is provided the cart will be associated with the first Region available. If no items are provided the cart will be empty after creation. If a user is logged in the cart's customer id and email will be set.

func (c *CreateCart) Create(config *medusa.Config) (*CreateResponse, error) {
	const path = `/store/carts`
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetData(c).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(CreateResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(CreateData)
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
