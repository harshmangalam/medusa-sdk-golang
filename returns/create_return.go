package returns

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type CreateReturn struct {
	OrderId        string `json:"order_id"`
	Items          []any  `json:"items"`
	ReturnShipping any    `json:"return_shipping"`
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

}
