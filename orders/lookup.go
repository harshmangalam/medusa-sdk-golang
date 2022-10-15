package orders

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

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

func (l *Lookup) Apply(cartId string, config *medusa.Config) (*Order, error) {
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

	respBody := new(ResponseBody)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}

	return respBody.Order, nil
}
