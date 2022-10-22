package schema

import "time"

type Order struct {
	CustomerId           string     `json:"customer_id"`
	Email                string     `json:"email"`
	RegionId             string     `json:"region_id"`
	CurrencyCode         string     `json:"currency_code"`
	Id                   string     `json:"id,omitempty"`
	Status               string     `json:"status,omitempty"`
	FulfillmentStatus    string     `json:"fulfillment_status,omitempty"`
	PaymentStatus        string     `json:"payment_status,omitempty"`
	DisplayId            string     `json:"display_id,omitempty"`
	CartId               string     `json:"cart_id,omitempty"`
	Cart                 any        `json:"cart,omitempty"`
	Customer             any        `json:"customer,omitempty"`
	BillingAddressId     string     `json:"billing_address_id,omitempty"`
	BillingAddress       any        `json:"billing_address,omitempty"`
	ShippingAddressId    string     `json:"shipping_address_id,omitempty"`
	ShippingAddress      any        `json:"shipping_address"`
	Region               any        `json:"region,omitempty"`
	Currency             any        `json:"currency,omitempty"`
	Tax_rate             int        `json:"tax_rate,omitempty"`
	Discounts            []any      `json:"discounts,omitempty"`
	Gift_cards           any        `json:"gift_cards,omitempty"`
	Payments             []any      `json:"payments,omitempty"`
	Fulfillments         []any      `json:"fulfillments,omitempty"`
	Returns              []any      `json:"returns,omitempty"`
	Claims               []any      `json:"claims,omitempty"`
	Refunds              []any      `json:"refunds,omitempty"`
	Swaps                []any      `json:"swaps,omitempty"`
	DraftOrderId         string     `json:"draft_order_id,omitempty"`
	DraftOrder           any        `json:"draft_order,omitempty"`
	Items                []any      `json:"items,omitempty"`
	Edits                []any      `json:"edits,omitempty"`
	GiftCardTransactions []any      `json:"gift_card_transactions,omitempty"`
	CanceledAt           *time.Time `json:"canceled_at,omitempty"`
	NoNotification       bool       `json:"no_notification,omitempty"`
	IdempotencyKey       string     `json:"idempotency_key,omitempty"`
	ExternalId           string     `json:"external_id,omitempty"`
	SalesChannelId       string     `json:"sales_channel_id,omitempty"`
	SalesChannel         any        `json:"sales_channel,omitempty"`
	ShippingTotal        int        `json:"shipping_total,omitempty"`
	DiscountTotal        int        `json:"discount_total,omitempty"`
	TaxTotal             int        `json:"tax_total,omitempty"`
	RefundedTotal        int        `json:"refunded_total,omitempty"`
	Total                int        `json:"total,omitempty"`
	Subtotal             int        `json:"subtotal,omitempty"`
	PaidTotal            int        `json:"paid_total,omitempty"`
	RefundableAmount     int        `json:"refundable_amount,omitempty"`
	GiftCardTotal        int        `json:"gift_card_total,omitempty"`
	GiftCardTaxTotal     int        `json:"gift_card_tax_total,omitempty"`
}
