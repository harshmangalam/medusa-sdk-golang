package customers

import "github.com/harshmngalam/medusa-sdk-golang/common"

type ListQuery struct {
	Q                 string                 `json:"q,omitempty" url:"q,omitempty"`
	Id                []string               `json:"id,omitempty" url:"id,omitempty"`
	CreatdAt          *common.DateComparison `json:"created_at,omitempty" url:"created_at,omitempty"`
	UpdatedAt         *common.DateComparison `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	CanceledAt        *common.DateComparison `json:"canceled_at,omitempty" url:"canceled_at,omitempty"`
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
}

func NewListQuery() *ListQuery {
	return new(ListQuery)
}
