package products

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

type ListVariantData struct {
	Variants []*schema.ProductVariant `json:"variants"`
}

type ListVariantResponse struct {
	// Success response
	Data *ListVariantData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type InventoryQuantity struct {
	Lt  int `json:"lt"`
	Gt  int `json:"Gt"`
	Lte int `json:"Lte"`
	Gte int `json:"Gte"`
}

type ListProductVariant struct {
	// A comma separated list of Product Variant ids to filter by.
	Ids string `json:"ids" url:"ids"`

	// A comma separated list of Product Variant relations to load.
	Expand string `json:"expand" url:"expand"`

	// How many product variants to skip in the result.
	Offset string `json:"offset" url:"offset"`

	// Maximum number of Product Variants to return.
	Limit string `json:"limit" url:"limit"`

	// product variant title to search for.
	Title any `json:"title" url:"title"`

	// Filter by available inventory quantity
	InventoryQuantity any `json:"inventory_quantity" url:"inventory_quantity"`
}

func NewListProuductVariant() *ListProductVariant {
	return new(ListProductVariant)
}

func (l *ListProductVariant) SetIds(ids string) *ListProductVariant {
	l.Ids = ids
	return l
}

func (l *ListProductVariant) SetExpand(expand string) *ListProductVariant {
	l.Expand = expand
	return l
}

func (l *ListProductVariant) SetOffset(offset string) *ListProductVariant {
	l.Offset = offset
	return l
}

func (l *ListProductVariant) SetLimit(limit string) *ListProductVariant {
	l.Limit = limit
	return l
}

func (l *ListProductVariant) SetTitle(title any) *ListProductVariant {
	l.Title = title
	return l
}

func (l *ListProductVariant) SetInventoryQuantity(invQty any) *ListProductVariant {
	l.InventoryQuantity = invQty
	return l
}

// Retrieves a list of Product Variants
func (l *ListProductVariant) List(config *medusa.Config) (*ListVariantResponse, error) {
	path := "/store/variants"

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

	respBody := new(ListVariantResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ListVariantData)
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
