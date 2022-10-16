package giftcards

import (
	"time"

	"github.com/harshmngalam/medusa-sdk-golang/orders"
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
