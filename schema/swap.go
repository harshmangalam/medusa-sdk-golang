package schema

import (
	"time"
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

// Swaps can be created when a Customer wishes to exchange Products that they have purchased to different Products. Swaps consist of a Return of previously purchased Products and a Fulfillment of new Products, the amount paid for the Products being returned will be used towards payment for the new Products. In the case where the amount paid for the the Products being returned exceed the amount to be paid for the new Products, a Refund will be issued for the difference.

type Swap struct {
	FulfillmentStatus FulfillmentStatus `json:"fulfillment_status"`
	PaymentStatus     PaymentStatus     `json:"payment_status"`
	OrderId           string            `json:"order_id"`
	Id                string            `json:"id"`
	Order             *Order            `json:"order"`
	AdditionalItems   []any             `json:"additional_items"`
	ReturnOrder       any               `json:"eturn_order"`
	Fulfillments      []any             `json:"fulfillments"`
	Payment           any               `json:"payment"`
	DifferenceDue     int               `json:"difference_due"`
	ShippingaddressId string            `json:"shipping_address_id"`
	ShippingAddress   *Address          `json:"shipping_address"`
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
