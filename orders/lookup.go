package orders

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

type LookupOrderData struct {
	Order *schema.Order `json:"order"`
}

type LookupOrderResponse struct {
	// Success response
	Data *LookupOrderData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type ShippingAddress struct {
	PostalCode string `json:"postal_code,omitempty" url:"postal_code,omitempty"`
}
type Lookup struct {
	DisplayId       string           `json:"display_id" url:"display_id"`
	Email           string           `json:"email" url:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty" url:"shipping_address,omitempty"`
}

func NewLookup() *Lookup {
	return new(Lookup)
}

func (l *Lookup) SetDisplayId(displayId string) *Lookup {
	l.DisplayId = displayId
	return l
}

func (l *Lookup) SetEmail(email string) *Lookup {
	l.Email = email
	return l
}

func (l *Lookup) SetShippingAddress(addr *ShippingAddress) *Lookup {
	l.ShippingAddress = addr
	return l
}

// Look up an order using filters.
func (l *Lookup) Lookup(cartId string, config *medusa.Config) (*LookupOrderResponse, error) {
	path := "/store/orders"
	qs, err := query.Values(l)
	if err != nil {
		return nil, err
	}

	parseStr := qs.Encode()

	path = fmt.Sprintf("%v?%v", path, parseStr)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(LookupOrderResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(LookupOrderData)
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
