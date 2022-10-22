package customers

import (
	"encoding/json"
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
	// Query used for searching orders.
	Q string `json:"q,omitempty" url:"q,omitempty"`

	// Id of the order to search for.
	Id []string `json:"id,omitempty" url:"id,omitempty"`

	// The offset in the resulting orders.
	Offset int `json:"offset" url:"offset"`

	// How many orders to return.
	Limit int `json:"limit" url:"limit"`

	// (Comma separated string) Which relations should be expanded in the resulting orders.
	Expand string `json:"expand,omitempty" url:"expand,omitempty"`

	// (Comma separated string) Which fields should be included in the resulting orders.
	Fields string `json:"fields,omitempty" url:"fields,omitempty"`

	// Status to search for.
	Status []string `json:"status,omitempty" url:"status,omitempty"`

	// Fulfillment status to search for.
	FulfillmentStatus []string `json:"fulfillment_status,omitempty" url:"fulfillment_status,omitempty"`

	// Payment status to search for.
	PaymentStatus []string `json:"payment_status,omitempty" url:"payment_status,omitempty"`

	// Display id to search for.
	DisplayId string `json:"display_id,omitempty" url:"display_id,omitempty"`

	// Cart id
	CartId string `json:"cart_id,omitempty" url:"cart_id,omitempty"`

	// Email
	Email string `json:"email,omitempty" url:"email,omitempty"`
	// Region id
	RegionId string `json:"region_id,omitempty" url:"region_id,omitempty"`

	// The 3 character ISO currency code to set prices based on.
	CurrencyCode string `json:"currency_code,omitempty" url:"currency_code,omitempty"`

	// Tax rate
	TaxRate string `json:"tax_rate,omitempty" url:"tax_rate,omitempty"`

	// Date comparison for when resulting collections were created.
	CreatedAt *common.DateComparison `json:"created_at,omitempty" url:"created_at,omitempty"`

	// Date comparison for when resulting collections were updated.
	UpdatedAt *common.DateComparison `json:"updated_at,omitempty" url:"updated_at,omitempty"`

	// Date comparison for when resulting collections were canceled.
	CanceledAt *common.DateComparison `json:"canceled_at,omitempty" url:"canceled_at,omitempty"`
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

// Retrieves a list of a Customer's Orders.
func (l *ListOrderQuery) List(config *medusa.Config) (*ListOrderResponse, error) {
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

	respBody := new(ListOrderResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ListOrderData)
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
