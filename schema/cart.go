package schema

import (
	"time"
)

type CartTypeEnum string

const (
	CartTypeDefault     CartTypeEnum = "default"
	CartTypeSwap        CartTypeEnum = "swap"
	CartTypeDraftOrder  CartTypeEnum = "draft_order"
	CartTypePaymentLink CartTypeEnum = "payment_link"
	CartTypeClaim       CartTypeEnum = "claim"
)

type Cart struct {
	// The cart's ID
	Id string `json:"id"`

	// The email associated with the cart
	Email string `json:"email"`

	// The billing address's ID
	BillingAddressId string `json:"billing_address_id"`

	// An address.
	BillingAddress *Address `json:"billing_address"`

	// The shipping address's ID
	ShippingAddressId string `json:"shipping_address_id"`

	// An address.
	ShippingAddress *Address `json:"shipping_address"`

	// Available if the relation items is expanded.
	Items []*LineItem `json:"items"`

	// The region's ID
	RegionId string `json:"region_id"`

	// A region object. Available if the relation region is expanded.
	Region any `json:"region"`

	// Available if the relation discounts is expanded.
	Discounts []any `json:"discounts"`

	// Available if the relation gift_cards is expanded
	GiftCards []*GiftCard `json:"gift_cards"`

	// The customer's ID
	CustomerId string `json:"customer_id"`

	// A customer object. Available if the relation customer is expanded.
	Customer *Customer `json:"customer"`

	// Payment Sessions are created when a Customer initilizes the checkout flow, and can be used to hold the state of a payment flow. Each Payment Session is controlled by a Payment Provider, who is responsible for the communication with external payment services. Authorized Payment Sessions will eventually get promoted to Payments to indicate that they are authorized for capture/refunds/etc.

	PaymentSession any `json:"payment_session"`

	// The payment sessions created on the cart.
	PaymentSessions []any `json:"payment_sessions"`

	// The payment's ID if available
	PaymentId string `json:"payment_id"`

	// Payments represent an amount authorized with a given payment method, Payments can be captured, canceled or refunded.
	Payment any `json:"payment"`

	// The shipping methods added to the cart.
	ShippingMethods any `json:"shipping_methods"`

	// The cart's type.
	Type CartTypeEnum `json:"type"`

	// The date with timezone at which the cart was completed.
	CompletedAt *time.Time `json:"completed_at"`

	// The date with timezone at which the payment was authorized.
	PaymentAuthorizedAt *time.Time `json:"payment_authorized_at"`

	// Randomly generated key used to continue the completion of a cart in case of failure.
	Idempotency_key string `json:"idempotency_key"`

	// Example: {"ip":"::1","user_agent":"PostmanRuntime/7.29.2"}
	// The context of the cart which can include info like IP or user agent.
	Context map[string]any `json:"context"`

	// The sales channel ID the cart is associated with.
	SalesChannelId string `json:"sales_channel_id"`

	// A sales channel object. Available if the relation sales_channel is expanded.
	SalesChannel any `json:"sales_channel"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// An optional key-value map with additional details
	Metadata any `json:"metadata,omitempty"`

	// The total of shipping
	ShippingTotal uint `json:"shipping_total"`

	// The total of discount
	DiscountTotal uint `json:"discount_total"`

	// The total of tax
	TaxTotal uint `json:"tax_total"`

	// The total amount refunded if the order associated with this cart is returned.
	RefundedTotal uint `json:"refunded_total"`

	// The total amount of the cart
	Total uint `json:"total"`

	// The subtotal of the cart
	Subtotal uint `json:"subtotal"`

	// The amount that can be refunded
	RefundableAmount uint `json:"refundable_amount"`

	// The total of gift cards
	GiftCardTaxTotal uint `json:"gift_card_tax_total"`

	// The total of gift cards with taxes
	GiftCardTotal uint `json:"gift_card_total"`
}
