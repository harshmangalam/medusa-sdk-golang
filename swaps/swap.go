package swaps

import (
	"time"

	"github.com/harshmngalam/medusa-sdk-golang/address"
	"github.com/harshmngalam/medusa-sdk-golang/orders"
)

type FulfillmentStatus string
type PaymentStatus string

const (
	FulfillmentNotFulfilled   FulfillmentStatus = "not_fulfilled"
	FulfillmentFulfilled      FulfillmentStatus = "fulfilled"
	FulfillmentShipped        FulfillmentStatus = "shipped"
	FulfillmentCanceled       FulfillmentStatus = "canceled"
	FulfillmentRequiresAction FulfillmentStatus = "requires_action"
)

const (
	PaymentNotPaid            PaymentStatus = "not_paid"
	PaymentAwaiting           PaymentStatus = "awaiting"
	PaymentCaptured           PaymentStatus = "captured"
	PaymentConfirmed          PaymentStatus = "confirmed"
	PaymentCanceled           PaymentStatus = "canceled"
	PaymentDifferenceRefunded PaymentStatus = "difference_refunded"
	PaymentPartiallyRefunded  PaymentStatus = "partially_refunded"
	PaymentRefunded           PaymentStatus = "refunded"
	RPaymentequiresAction     PaymentStatus = "requires_action"
)

type Swap struct {
	FulfillmentStatus FulfillmentStatus `json:"fulfillment_status"`
	PaymentStatus     PaymentStatus     `json:"payment_status"`
	OrderId           string            `json:"order_id"`
	Id                string            `json:"id"`
	Order             *orders.Order     `json:"order"`
	AdditionalItems   []any             `json:"additional_items"`
	ReturnOrder       any               `json:"eturn_order"`
	Fulfillments      []any             `json:"fulfillments"`
	Payment           any               `json:"payment"`
	DifferenceDue     int               `json:"difference_due"`
	ShippingaddressId string            `json:"shipping_address_id"`
	ShippingAddress   *address.Address  `json:"shipping_address"`
	ShippingMethods   []any             `json:"shipping_methods"`
	CartId            string            `json:"cart_id"`
	Cart              any               `json:"cart"`
	AllowBackorder    bool              `json:"allow_backorder"`
	IdempotencyKey    string            `json:"idempotency_key"`
	ConfirmedAt       *time.Time        `json:"confirmed_at"`
	CanceledAt        *time.Time        `json:"canceled_at"`
	NoNotification    bool              `json:"no_notification"`
	CreatedAt         *time.Time        `json:"created_at"`
	UpdatedAt         *time.Time        `json:"updated_at"`
	DeletedAt         *time.Time        `json:"deleted_at"`
	Metadata          map[string]any    `json:"metadata"`
}

type ResponseBody struct {
	Swap *Swap `json:"swap"`
}
