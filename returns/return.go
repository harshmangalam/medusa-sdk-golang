package returns

import (
	"encoding/json"
	"net/http"
	"time"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/orders"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/swaps"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type Return struct {
	RefundAmount   int            `json:"refund_amount"`
	Id             string         `json:"id"`
	Status         any            `json:"status"`
	Items          []any          `json:"items"`
	SwapId         string         `json:"swap_id"`
	Swap           *swaps.Swap    `json:"swap"`
	OrderId        string         `json:"order_id"`
	Order          *orders.Order  `json:"order"`
	ClaimOrderId   string         `json:"claim_order_id"`
	ClaimOrder     any            `json:"claim_order"`
	ShippingMethod any            `json:"shipping_method"`
	ShippingData   any            `json:"shipping_data"`
	NoNotification bool           `json:"no_notification"`
	IdempotencyKey string         `json:"idempotency_key"`
	ReceivedAt     *time.Time     `json:"received_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      *time.Time     `json:"deleted_at"`
	Metadata       map[string]any `json:"metadata"`
}

type CreateReturn struct {
	OrderId        string `json:"order_id"`
	Items          []any  `json:"items"`
	ReturnShipping any    `json:"return_shipping"`
}

type ReturnResponse struct {
	Return *Return `json:"return"`
}

func NewCreateRetun() *CreateReturn {
	return new(CreateReturn)
}

func (c *CreateReturn) SetOrderId(orderId string) *CreateReturn {
	c.OrderId = orderId
	return c
}
func (c *CreateReturn) SetItems(items []any) *CreateReturn {
	c.Items = items
	return c

}

func (c *CreateReturn) SetReturnShipping(shipping any) *CreateReturn {
	c.ReturnShipping = shipping
	return c
}

func (c *CreateReturn) Create(config *medusa.Config) (*Return, error) {
	path := "/store/returns"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(c).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(ReturnResponse)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}

	return respBody.Return, nil
}