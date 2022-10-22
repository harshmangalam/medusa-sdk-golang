package schema

import (
	"time"
)

type Return struct {
	RefundAmount   int            `json:"refund_amount"`
	Id             string         `json:"id"`
	Status         any            `json:"status"`
	Items          []any          `json:"items"`
	SwapId         string         `json:"swap_id"`
	Swap           *Swap          `json:"swap"`
	OrderId        string         `json:"order_id"`
	Order          *Order         `json:"order"`
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
