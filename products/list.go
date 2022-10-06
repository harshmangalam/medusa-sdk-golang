package products

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/common"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ListQuery struct {
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
	CreatdAt *common.DateComparison `json:"created_at,omitempty" url:"created_at,omitempty"`

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

// create new product query
func NewListQuery() *ListQuery {
	p := new(ListQuery)
	p.Offset = 0
	p.Limit = 100
	return p
}

// set product query
func (p *ListQuery) SetQ(q string) *ListQuery {
	p.Q = q
	return p
}

// set product ids
func (p *ListQuery) SetIds(ids []string) *ListQuery {
	p.Ids = ids
	return p
}

// set collection ids
func (p *ListQuery) SetCollectionIds(collectionIds []string) *ListQuery {
	p.CollectionIds = collectionIds
	return p
}

// set tags
func (p *ListQuery) SetTags(tags []string) *ListQuery {
	p.Tags = tags
	return p
}

// set title
func (p *ListQuery) SetTitle(title string) *ListQuery {
	p.Title = title
	return p
}

// set description
func (p *ListQuery) SetDescription(description string) *ListQuery {
	p.Description = description
	return p
}

// set handle
func (p *ListQuery) SetHandle(handle string) *ListQuery {
	p.Handle = handle
	return p
}

// set gift card
func (p *ListQuery) SetIsGiftcard(isGiftcard bool) *ListQuery {
	p.IsGiftcard = isGiftcard
	return p
}

// set type

func (p *ListQuery) SetType(productType string) *ListQuery {
	p.Type = productType
	return p
}

func (p *ListQuery) SetCreatedAt(creatdAt *common.DateComparison) *ListQuery {
	p.CreatdAt = creatdAt
	return p
}

func (p *ListQuery) SetUpdatedAt(updatedAt *common.DateComparison) *ListQuery {
	p.UpdatedAt = updatedAt
	return p
}

func (p *ListQuery) SetOffset(offset int) *ListQuery {
	p.Offset = offset
	return p
}

func (p *ListQuery) SetLimit(limit int) *ListQuery {
	p.Limit = limit
	return p
}

func (p *ListQuery) SetExpand(expand string) *ListQuery {
	p.Expand = expand
	return p
}

func (p *ListQuery) SetFields(fields string) *ListQuery {
	p.Fields = fields
	return p
}

// Retrieve a list of Products.
func (c *ListQuery) List(config *medusa.Config) ([]byte, error) {
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
	return body, nil
}