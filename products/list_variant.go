package products

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

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

func (l *VariantListQuery) List(config *medusa.Config) ([]byte, error) {
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

	return body, nil
}
