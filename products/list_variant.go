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

type VariantListQuery struct {
	Ids               string `json:"ids" url:"ids"`
	Expand            string `json:"expand" url:"expand"`
	Offset            string `json:"offset" url:"offset"`
	Limit             string `json:"limit" url:"limit"`
	Title             any    `json:"title" url:"title"`
	InventoryQuantity any    `json:"inventory_quantity" url:"inventory_quantity"`
}

func NewVariantListQuery() *VariantListQuery {
	return new(VariantListQuery)
}

func (l *VariantListQuery) SetIds(ids string) *VariantListQuery {
	l.Ids = ids
	return l
}

func (l *VariantListQuery) SetExpand(expand string) *VariantListQuery {
	l.Expand = expand
	return l
}

func (l *VariantListQuery) SetOffset(offset string) *VariantListQuery {
	l.Offset = offset
	return l
}

func (l *VariantListQuery) SetLimit(limit string) *VariantListQuery {
	l.Limit = limit
	return l
}

func (l *VariantListQuery) SetTitle(title any) *VariantListQuery {
	l.Title = title
	return l
}

func (l *VariantListQuery) SetInventoryQuantity(invQty any) *VariantListQuery {
	l.InventoryQuantity = invQty
	return l
}

func (l *VariantListQuery) List(config *medusa.Config) (*ListVariantResponse, error) {
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
