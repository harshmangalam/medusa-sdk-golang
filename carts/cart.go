package carts

import (
	"time"

	"github.com/harshmngalam/medusa-sdk-golang/address"
	"github.com/harshmngalam/medusa-sdk-golang/customers"
	giftcards "github.com/harshmngalam/medusa-sdk-golang/gift_cards"
)

type Cart struct {
	Id                  string                `json:"id"`
	Email               string                `json:"email"`
	BillingddressId     string                `json:"billing_address_id"`
	BillingAddress      *address.Address      `json:"billing_address"`
	ShippingAddressId   string                `json:"shipping_address_id"`
	ShippingAddress     *address.Address      `json:"shipping_address"`
	Items               []any                 `json:"items"`
	RegionId            string                `json:"region_id"`
	Region              any                   `json:"region"`
	Discounts           []any                 `json:"discounts"`
	GiftCards           []*giftcards.GiftCard `json:"gift_cards"`
	CustomerId          string                `json:"customer_id"`
	Customer            *customers.Customer   `json:"customer"`
	PaymentSession      any                   `json:"payment_session"`
	PaymentSessions     []any                 `json:"payment_sessions"`
	PaymentId           string                `json:"payment_id"`
	Payment             any                   `json:"payment"`
	ShippingMethods     any                   `json:"shipping_methods"`
	Type                any                   `json:"type"`
	CompletedAt         *time.Time            `json:"completed_at"`
	PaymentAuthorizedAt *time.Time            `json:"payment_authorized_at"`
	Idempotency_key     string                `json:"idempotency_key"`
	Context             map[string]any        `json:"context"`
	SalesChannelId      string                `json:"sales_channel_id"`
	SalesChannel        any                   `json:"sales_channel"`
	CreatedAt           *time.Time            `json:"created_at,omitempty"`
	UpdatedAt           *time.Time            `json:"updated_at,omitempty"`
	DeletedAt           *time.Time            `json:"deleted_at,omitempty"`
	Metadata            any                   `json:"metadata,omitempty"`
	ShippingTotal       uint                  `json:"shipping_total"`
	DiscountTotal       uint                  `json:"discount_total"`
	TaxTotal            uint                  `json:"tax_total"`
	RefundedTotal       uint                  `json:"refunded_total"`
	Total               uint                  `json:"total"`
	Subtotal            uint                  `json:"subtotal"`
	RefundableAmount    uint                  `json:"refundable_amount"`
	GiftCardTaxTotal    uint                  `json:"gift_card_tax_total"`
	GiftCardTotal       uint                  `json:"gift_card_total"`
}
