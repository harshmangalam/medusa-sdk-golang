package customers

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/common"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ListOrderData struct {
	// Array of orders
	Orders []*schema.Order `json:"orders"`

	// The total number of items available
	Count uint `json:"count"`

	// The number of items skipped before these items
	Offset uint `json:"offset"`

	// The number of items per page
	Limit uint `json:"limit"`
}

type ListOrderResponse struct {
	// Success response
	Data *ListOrderData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type ListOrderQuery struct {
	Q                 string                 `json:"q,omitempty" url:"q,omitempty"`
	Id                []string               `json:"id,omitempty" url:"id,omitempty"`
	Offset            int                    `json:"offset" url:"offset"`
	Limit             int                    `json:"limit" url:"limit"`
	Expand            string                 `json:"expand,omitempty" url:"expand,omitempty"`
	Fields            string                 `json:"fields,omitempty" url:"fields,omitempty"`
	Status            []string               `json:"status,omitempty" url:"status,omitempty"`
	FulfillmentStatus []string               `json:"fulfillment_status,omitempty" url:"fulfillment_status,omitempty"`
	PaymentStatus     []string               `json:"payment_status,omitempty" url:"payment_status,omitempty"`
	CartId            string                 `json:"cart_id,omitempty" url:"cart_id,omitempty"`
	Email             string                 `json:"email,omitempty" url:"email,omitempty"`
	RegionId          string                 `json:"region_id,omitempty" url:"region_id,omitempty"`
	CurrencyCode      string                 `json:"currency_code,omitempty" url:"currency_code,omitempty"`
	TaxRate           string                 `json:"tax_rate,omitempty" url:"tax_rate,omitempty"`
	CreatedAt         *common.DateComparison `json:"created_at,omitempty" url:"created_at,omitempty"`
	UpdatedAt         *common.DateComparison `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	CanceledAt        *common.DateComparison `json:"canceled_at,omitempty" url:"canceled_at,omitempty"`
}

func NewListOrderQuery() *ListOrderQuery {
	l := new(ListOrderQuery)
	l.Limit = 10
	l.Offset = 0
	return l
}

func (l *ListOrderQuery) SetQ(q string) *ListOrderQuery {
	l.Q = q
	return l
}
func (l *ListOrderQuery) SetId(id []string) *ListOrderQuery {
	l.Id = id
	return l
}

func (l *ListOrderQuery) SetOffset(offset int) *ListOrderQuery {
	l.Offset = offset
	return l
}

func (l *ListOrderQuery) SetLimit(limit int) *ListOrderQuery {
	l.Offset = limit
	return l
}

func (l *ListOrderQuery) SetExpand(expand string) *ListOrderQuery {
	l.Expand = expand
	return l
}

func (l *ListOrderQuery) SetFields(fields string) *ListOrderQuery {
	l.Fields = fields
	return l
}

func (l *ListOrderQuery) SetStatus(status []string) *ListOrderQuery {
	l.Status = status
	return l
}

func (l *ListOrderQuery) SetFulfillmentStatus(fulfillmentStatus []string) *ListOrderQuery {
	l.FulfillmentStatus = fulfillmentStatus
	return l
}

func (l *ListOrderQuery) SetPaymentStatus(paymentStatus []string) *ListOrderQuery {
	l.PaymentStatus = paymentStatus
	return l
}

func (l *ListOrderQuery) SetCartId(cartId string) *ListOrderQuery {
	l.CartId = cartId
	return l
}

func (l *ListOrderQuery) SetEmail(email string) *ListOrderQuery {
	l.Email = email
	return l
}

func (l *ListOrderQuery) SetRegionId(regionId string) *ListOrderQuery {
	l.RegionId = regionId
	return l
}

func (l *ListOrderQuery) SetCurrencyCode(currencyCode string) *ListOrderQuery {
	l.CurrencyCode = currencyCode
	return l
}

func (l *ListOrderQuery) SetTaxRate(taxRate string) *ListOrderQuery {
	l.TaxRate = taxRate
	return l
}

func (l *ListOrderQuery) SetCreatedAt(createdAt *common.DateComparison) *ListOrderQuery {
	l.CreatedAt = createdAt
	return l
}

func (l *ListOrderQuery) SetUpdatedAt(updatedAt *common.DateComparison) *ListOrderQuery {
	l.UpdatedAt = updatedAt
	return l
}

func (l *ListOrderQuery) SetCancledAt(canceledAt *common.DateComparison) *ListOrderQuery {
	l.CanceledAt = canceledAt
	return l
}

func (l *ListOrderQuery) Apply(config *medusa.Config) ([]byte, error) {
	path := "/store/customers/me/orders"

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
