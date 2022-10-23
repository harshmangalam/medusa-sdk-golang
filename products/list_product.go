package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmangalam/medusa-sdk-golang"
	"github.com/harshmangalam/medusa-sdk-golang/common"
	"github.com/harshmangalam/medusa-sdk-golang/request"
	"github.com/harshmangalam/medusa-sdk-golang/response"
	"github.com/harshmangalam/medusa-sdk-golang/schema"
	"github.com/harshmangalam/medusa-sdk-golang/utils"
)

type ListProductData struct {
	// Array of product
	Products []*schema.Product `json:"products"`

	// The total number of items available
	Count uint `json:"count"`

	// The number of items skipped before these items
	Offset uint `json:"offset"`

	// The number of items per page
	Limit uint `json:"limit"`
}

type ListProductResponse struct {
	// Success response
	Data *ListProductData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type ListProduct struct {
	// Query used for searching products by title, description, variant's title, variant's sku, and collection's title
	Q string `json:"q,omitempty" url:"q,omitempty"`

	// product IDs to search for.
	Ids []string `json:"id,omitempty" url:"id,omitempty"`

	// Collection IDs to search for
	CollectionIds []string `json:"collection_id" url:"collection_id"`

	// Tag IDs to search for
	Tags []string `json:"tags,omitempty" url:"tags,omitempty"`

	// title to search for.
	Title string `json:"title,omitempty" url:"title,omitempty"`

	// description to search for
	Description string `json:"description,omitempty" url:"description,omitempty"`

	// handle to search for.
	Handle string `json:"handle,omitempty" url:"handle,omitempty"`

	// Search for giftcards using is_giftcard=true.
	IsGiftcard bool `json:"is_giftcard,omitempty" url:"is_giftcard,omitempty"`

	// type to search for.
	Type string `json:"type,omitempty" url:"type,omitempty"`

	// Date comparison for when resulting products were created.
	CreatedAt *common.DateComparison `json:"created_at,omitempty" url:"created_at,omitempty"`

	// Date comparison for when resulting products were updated.
	UpdatedAt *common.DateComparison `json:"updated_at,omitempty" url:"updated_at,omitempty"`

	// How many products to skip in the result.
	Offset int `json:"offset" url:"offset"`

	// Limit the number of products returned.
	Limit int `json:"limit" url:"limit"`

	// (Comma separated) Which fields should be expanded in each order of the result.)
	Expand string `json:"expand,omitempty" url:"expand,omitempty"`

	// (Comma separated) Which fields should be included in each order of the result.
	Fields string `json:"fields,omitempty" url:"fields,omitempty"`
}

func NewListProduct() *ListProduct {
	p := new(ListProduct)
	p.Offset = 0
	p.Limit = 100
	return p
}

func (p *ListProduct) SetQ(q string) *ListProduct {
	p.Q = q
	return p
}

func (p *ListProduct) SetIds(ids []string) *ListProduct {
	p.Ids = ids
	return p
}

func (p *ListProduct) SetCollectionIds(collectionIds []string) *ListProduct {
	p.CollectionIds = collectionIds
	return p
}

func (p *ListProduct) SetTags(tags []string) *ListProduct {
	p.Tags = tags
	return p
}

func (p *ListProduct) SetTitle(title string) *ListProduct {
	p.Title = title
	return p
}

func (p *ListProduct) SetDescription(description string) *ListProduct {
	p.Description = description
	return p
}

func (p *ListProduct) SetHandle(handle string) *ListProduct {
	p.Handle = handle
	return p
}

func (p *ListProduct) SetIsGiftcard(isGiftcard bool) *ListProduct {
	p.IsGiftcard = isGiftcard
	return p
}

func (p *ListProduct) SetType(productType string) *ListProduct {
	p.Type = productType
	return p
}

func (p *ListProduct) SetCreatedAt(creatdAt *common.DateComparison) *ListProduct {
	p.CreatedAt = creatdAt
	return p
}

func (p *ListProduct) SetUpdatedAt(updatedAt *common.DateComparison) *ListProduct {
	p.UpdatedAt = updatedAt
	return p
}

func (p *ListProduct) SetOffset(offset int) *ListProduct {
	p.Offset = offset
	return p
}

func (p *ListProduct) SetLimit(limit int) *ListProduct {
	p.Limit = limit
	return p
}

func (p *ListProduct) SetExpand(expand string) *ListProduct {
	p.Expand = expand
	return p
}

func (p *ListProduct) SetFields(fields string) *ListProduct {
	p.Fields = fields
	return p
}

// Retrieve a list of Products.
func (c *ListProduct) List(config *medusa.Config) (*ListProductResponse, error) {
	path := "/store/products"

	qs, err := query.Values(c)
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
	respBody := new(ListProductResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ListProductData)
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
