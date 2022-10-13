package customers

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

func NewListQuery() *ListQuery {
	return new(ListQuery)
}

func (l *ListQuery) SetQ(q string) *ListQuery {
	l.Q = q
	return l
}
func (l *ListQuery) SetId(id []string) *ListQuery {
	l.Id = id
	return l
}

func (l *ListQuery) SetOffset(offset int) *ListQuery {
	l.Offset = offset
	return l
}

func (l *ListQuery) SetLimit(limit int) *ListQuery {
	l.Offset = limit
	return l
}

func (l *ListQuery) SetExpand(expand string) *ListQuery {
	l.Expand = expand
	return l
}

func (l *ListQuery) SetFields(fields string) *ListQuery {
	l.Fields = fields
	return l
}

func (l *ListQuery) SetStatus(status []string) *ListQuery {
	l.Status = status
	return l
}

func (l *ListQuery) SetFulfillmentStatus(fulfillmentStatus []string) *ListQuery {
	l.FulfillmentStatus = fulfillmentStatus
	return l
}

func (l *ListQuery) SetPaymentStatus(paymentStatus []string) *ListQuery {
	l.PaymentStatus = paymentStatus
	return l
}

func (l *ListQuery) SetCartId(cartId string) *ListQuery {
	l.CartId = cartId
	return l
}

func (l *ListQuery) SetEmail(email string) *ListQuery {
	l.Email = email
	return l
}

func (l *ListQuery) SetRegionId(regionId string) *ListQuery {
	l.RegionId = regionId
	return l
}

func (l *ListQuery) SetCurrencyCode(currencyCode string) *ListQuery {
	l.CurrencyCode = currencyCode
	return l
}

func (l *ListQuery) SetTaxRate(taxRate string) *ListQuery {
	l.TaxRate = taxRate
	return l
}

func (l *ListQuery) SetCreatedAt(createdAt *common.DateComparison) *ListQuery {
	l.CreatedAt = createdAt
	return l
}

func (l *ListQuery) SetUpdatedAt(updatedAt *common.DateComparison) *ListQuery {
	l.UpdatedAt = updatedAt
	return l
}

func (l *ListQuery) SetCancledAt(canceledAt *common.DateComparison) *ListQuery {
	l.CanceledAt = canceledAt
	return l
}

func (l *ListQuery) ListOrders(config *medusa.Config) ([]byte, error) {
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
