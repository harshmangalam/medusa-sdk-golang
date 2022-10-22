package shippingoptions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ListShippingOptionData struct {
	ShippingOptions []*schema.ShippingOption `json:"shipping_options"`
}

type ListShippingOptionResponse struct {
	// Success response
	Data *ListShippingOptionData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type ListCartOptionsQuery struct {
	// Whether return Shipping Options should be included. By default all Shipping Options are returned.
	IsReturn bool `json:"is_return,omitempty" url:"is_return,omitempty"`

	// A comma separated list of Product ids to filter Shipping Options by.
	ProductIds string `json:"product_ids,omitempty" url:"product_ids,omitempty"`

	// the Region to retrieve Shipping Options from.
	RegionId string `json:"region_id,omitempty" url:"region_id,omitempty"`
}

func NewListCartOptionsQuery() *ListCartOptionsQuery {
	return new(ListCartOptionsQuery)
}

func (l *ListCartOptionsQuery) SetIsreturn(isReturn bool) *ListCartOptionsQuery {
	l.IsReturn = isReturn
	return l
}

func (l *ListCartOptionsQuery) SetProductIds(productIds string) *ListCartOptionsQuery {
	l.ProductIds = productIds
	return l
}

func (l *ListCartOptionsQuery) SetRegionId(regionId string) *ListCartOptionsQuery {
	l.RegionId = regionId
	return l
}

//Retrieves a list of Shipping Options.

func (l *ListCartOptionsQuery) List(config *medusa.Config) (*ListShippingOptionResponse, error) {
	path := "/store/shipping-options"

	qs, err := query.Values(l)
	if err != nil {
		return nil, err
	}

	parseStr := qs.Encode()

	path = fmt.Sprintf("%v?%v", path, parseStr)

	resp, err := request.
		NewRequest().
		SetMethod(http.MethodGet).
		SetPath(path).
		Send(config)

	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(ListShippingOptionResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ListShippingOptionData)
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
