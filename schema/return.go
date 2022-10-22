package schema

import (
	"time"
)

// Return orders hold information about Line Items that a Customer wishes to send back, along with how the items will be returned. Returns can be used as part of a Swap.

type Return struct {
	// The amount that should be refunded as a result of the return.
	RefundAmount int `json:"refund_amount"`

	// The return's ID
	Id string `json:"id"`

	// Status of the Return.
	Status any `json:"status"`

	// The Return Items that will be shipped back to the warehouse. Available if the relation items is expanded.
	Items []any `json:"items"`

	// The ID of the Swap that the Return is a part of.
	SwapId string `json:"swap_id"`

	// A swap object. Available if the relation swap is expanded.
	Swap *Swap `json:"swap"`

	// The ID of the Order that the Return is made from.
	OrderId string `json:"order_id"`

	// An order object. Available if the relation order is expanded.
	Order *Order `json:"order"`

	// The ID of the Claim that the Return is a part of.
	ClaimOrderId string `json:"claim_order_id"`

	// A claim order object. Available if the relation claim_order is expanded.
	ClaimOrder any `json:"claim_order"`

	// The Shipping Method that will be used to send the Return back. Can be null if the Customer facilitates the return shipment themselves. Available if the relation shipping_method is expanded.
	ShippingMethod any `json:"shipping_method"`

	// Data about the return shipment as provided by the Fulfilment Provider that handles the return shipment.
	ShippingData any `json:"shipping_data"`

	// When set to true, no notification will be sent related to this return.
	NoNotification bool `json:"no_notification"`

	// Randomly generated key used to continue the completion of the return in case of failure.
	IdempotencyKey string `json:"idempotency_key"`

	// The date with timezone at which the return was received.
	ReceivedAt *time.Time `json:"received_at"`

	// The date with timezone at which the resource was created.
	UpdatedAt *time.Time `json:"updated_at"`

	// The date with timezone at which the resource was updated.
	DeletedAt *time.Time `json:"deleted_at"`

	// An optional key-value map with additional details
	Metadata map[string]any `json:"metadata"`
}
