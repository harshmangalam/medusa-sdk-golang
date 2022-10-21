package carts

import (
	"time"

	"github.com/harshmngalam/medusa-sdk-golang/address"
	"github.com/harshmngalam/medusa-sdk-golang/customers"
	giftcards "github.com/harshmngalam/medusa-sdk-golang/gift_cards"
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
	BillingddressId string `json:"billing_address_id"`

	// An address.
	BillingAddress *address.Address `json:"billing_address"`

	// The shipping address's ID
	ShippingAddressId string `json:"shipping_address_id"`

	// An address.
	ShippingAddress *address.Address `json:"shipping_address"`

	// Available if the relation items is expanded.
	Items []any `json:"items"`

	// The region's ID
	RegionId string `json:"region_id"`

	// A region object. Available if the relation region is expanded.
	Region any `json:"region"`

	// Available if the relation discounts is expanded.
	Discounts []any `json:"discounts"`

	// Available if the relation gift_cards is expanded
	GiftCards []*giftcards.GiftCard `json:"gift_cards"`

	// The customer's ID
	CustomerId string `json:"customer_id"`

	// A customer object. Available if the relation customer is expanded.
	Customer *customers.Customer `json:"customer"`

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
	Type                any            `json:"type"`
	CompletedAt         *time.Time     `json:"completed_at"`
	PaymentAuthorizedAt *time.Time     `json:"payment_authorized_at"`
	Idempotency_key     string         `json:"idempotency_key"`
	Context             map[string]any `json:"context"`
	SalesChannelId      string         `json:"sales_channel_id"`
	SalesChannel        any            `json:"sales_channel"`
	CreatedAt           *time.Time     `json:"created_at,omitempty"`
	UpdatedAt           *time.Time     `json:"updated_at,omitempty"`
	DeletedAt           *time.Time     `json:"deleted_at,omitempty"`
	Metadata            any            `json:"metadata,omitempty"`
	ShippingTotal       uint           `json:"shipping_total"`
	DiscountTotal       uint           `json:"discount_total"`
	TaxTotal            uint           `json:"tax_total"`
	RefundedTotal       uint           `json:"refunded_total"`
	Total               uint           `json:"total"`
	Subtotal            uint           `json:"subtotal"`
	RefundableAmount    uint           `json:"refundable_amount"`
	GiftCardTaxTotal    uint           `json:"gift_card_tax_total"`
	GiftCardTotal       uint           `json:"gift_card_total"`
}
