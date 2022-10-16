package giftcards

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/orders"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type GiftCard struct {
	Code       string         `json:"code"`
	Value      int            `json:"value"`
	Balance    int            `json:"balance"`
	RegionId   string         `json:"region_id"`
	Id         string         `json:"id,omitempty"`
	Region     string         `json:"region,omitempty"`
	OrderId    string         `json:"order_id,omitempty"`
	Order      *orders.Order  `json:"order,omitempty"`
	IsDisabled bool           `json:"is_disabled,omitempty"`
	EndsAt     *time.Time     `json:"ends_at,omitempty"`
	CreatedAt  *time.Time     `json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
	DeletedAt  *time.Time     `json:"deleted_at,omitempty"`
	Metadata   map[string]any `json:"metadata,omitempty"`
}

type ResponseBody struct {
	GiftCard *GiftCard `json:"gift_card"`
}

func Retrieve(code string, config *medusa.Config) (*GiftCard, error) {
	path := fmt.Sprintf("/store/gift-cards/%v", code)
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

	return respBody.GiftCard, nil
}
