package orders

import "time"

type Order struct {
	CustomerId           string     `json:"customer_id"`
	Email                string     `json:"email"`
	RegionId             string     `json:"region_id"`
	CurrencyCode         string     `json:"currency_code"`
	Id                   string     `json:"id"`
	Status               string     `json:"status"`
	FulfillmentStatus    string     `json:"fulfillment_status"`
	PaymentStatus        string     `json:"payment_status"`
	DisplayId            string     `json:"display_id"`
	CartId               string     `json:"cart_id"`
	Cart                 any        `json:"cart"`
	Customer             any        `json:"customer"`
	BillingAddressId     string     `json:"billing_address_id"`
	BillingAddress       any        `json:"billing_address"`
	ShippingAddressId    string     `json:"shipping_address_id"`
	ShippingAddress      any        `json:"shipping_address_id"`
	Region               any        `json:"region"`
	Currency             any        `json:"currency"`
	Tax_rate             int        `json:"tax_rate"`
	Discounts            []any      `json:"discounts"`
	Gift_cards           any        `json:"gift_cards"`
	Payments             []any      `json:"payments"`
	Fulfillments         []any      `json:"fulfillments"`
	Returns              []any      `json:"returns"`
	Claims               []any      `json:"claims"`
	Refunds              []any      `json:"refunds"`
	Swaps                []any      `json:"swaps"`
	DraftOrderId         string     `json:"draft_order_id"`
	DraftOrder           any        `json:"draft_order"`
	Items                []any      `json:"items"`
	Edits                []any      `json:"edits"`
	GiftCardTransactions []any      `json:"gift_card_transactions"`
	CanceledAt           *time.Time `json:"canceled_at"`
	NoNotification       bool       `json:"no_notification"`
	IdempotencyKey       string     `json:"idempotency_key"`
	ExternalId           string     `json:"external_id"`
	SalesChannelId       string     `json:"sales_channel_id"`
	SalesChannel         any        `json:"sales_channel"`
	ShippingTotal        int        `json:"shipping_total"`
	DiscountTotal        int        `json:"discount_total"`
	TaxTotal             int        `json:"tax_total"`
	RefundedTotal        int        `json:"refunded_total"`
	Total                int        `json:"total"`
	Subtotal             int        `json:"subtotal"`
	PaidTotal            int        `json:"paid_total"`
	RefundableAmount     int        `json:"refundable_amount"`
	GiftCardTotal        int        `json:"gift_card_total"`
	GiftCardTaxTotal     int        `json:"gift_card_tax_total"`
}
