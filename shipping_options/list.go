package shippingoptions

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ListCartOptionsQuery struct {
	Isreturn   bool   `json:"is_return,omitempty" url:"is_return,omitempty"`
	ProductIds string `json:"product_ids,omitempty" url:"product_ids,omitempty"`
	RegionId   string `json:"region_id,omitempty" url:"region_id,omitempty"`
}

func NewListCartOptionsQuery() *ListCartOptionsQuery {
	return new(ListCartOptionsQuery)
}

func (l *ListCartOptionsQuery) SetIsreturn(isReturn bool) *ListCartOptionsQuery {
	l.Isreturn = isReturn
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

func (l *ListCartOptionsQuery) List(config *medusa.Config) ([]byte, error) {
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
	return body, nil
}
