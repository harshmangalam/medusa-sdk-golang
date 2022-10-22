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
	// The status of the Fulfillment of the Swap.
	FulfillmentStatus FulfillmentStatus `json:"fulfillment_status"`

	// The status of the Payment of the Swap. The payment may either refer to the refund of an amount or the authorization of a new amount.
	PaymentStatus PaymentStatus `json:"payment_status"`

	// The ID of the Order where the Line Items to be returned where purchased.
	OrderId string `json:"order_id"`

	// The swap's ID
	Id string `json:"id"`

	// An order object. Available if the relation order is expanded.
	Order *Order `json:"order"`

	// The new Line Items to ship to the Customer. Available if the relation additional_items is expanded.
	AdditionalItems []any `json:"additional_items"`

	// A return order object. The Return that is issued for the return part of the Swap. Available if the relation return_order is expanded.
	ReturnOrder any `json:"eturn_order"`

	// The Fulfillments used to send the new Line Items. Available if the relation fulfillments is expanded.
	Fulfillments []any `json:"fulfillments"`

	// Payments represent an amount authorized with a given payment method, Payments can be captured, canceled or refunded.
	Payment any `json:"payment"`

	// The difference that is paid or refunded as a result of the Swap. May be negative when the amount paid for the returned items exceed the total of the new Products.
	DifferenceDue int `json:"difference_due"`

	// The Address to send the new Line Items to - in most cases this will be the same as the shipping address on the Order.
	ShippingaddressId string `json:"shipping_address_id"`

	// An address
	ShippingAddress *Address `json:"shipping_address"`

	// The Shipping Methods used to fulfill the additional items purchased. Available if the relation shipping_methods is expanded.
	ShippingMethods []any `json:"shipping_methods"`

	// The id of the Cart that the Customer will use to confirm the Swap.
	CartId string `json:"cart_id"`

	// A cart object. Available if the relation cart is expanded.
	Cart any `json:"cart"`

	// If true, swaps can be completed with items out of stock
	AllowBackorder bool `json:"allow_backorder"`

	// Randomly generated key used to continue the completion of the swap in case of failure.
	IdempotencyKey string `json:"idempotency_key"`

	// The date with timezone at which the Swap was confirmed by the Customer.
	ConfirmedAt *time.Time `json:"confirmed_at"`

	// The date with timezone at which the Swap was canceled
	CanceledAt *time.Time `json:"canceled_at"`

	// If set to true, no notification will be sent related to this swap
	NoNotification bool `json:"no_notification"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at"`

	// An optional key-value map with additional details
	Metadata map[string]any `json:"metadata"`
}
