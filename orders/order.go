package orders

import "time"

type Order struct {
	CustomerId             string     `json:"customer_id"`
	Email                  string     `json:"email"`
	RegionId               string     `json:"region_id"`
	currency_code          string     `json:"currency_code"`
	Id                     string     `json:"id"`
	status                 string     `json:"status"`
	fulfillment_status     string     `json:"fulfillment_status"`
	payment_status         string     `json:"payment_status"`
	display_id             string     `json:"display_id"`
	cart_id                string     `json:"cart_id"`
	Cart                   any        `json:"cart"`
	Customer               any        `json:"customer"`
	billing_address_id     string     `json:"billing_address_id"`
	billing_address        any        `json:"billing_address"`
	shipping_address_id    string     `json:"shipping_address_id"`
	shipping_address       any        `json:"shipping_address_id"`
	Region                 any        `json:"region"`
	Currency               any        `json:"currency"`
	tax_rate               int        `json:"tax_rate"`
	Discounts              []any      `json:"discounts"`
	gift_cards             any        `json:"gift_cards"`
	payments               []any      `json:"payments"`
	fulfillments           []any      `json:"fulfillments"`
	Returns                []any      `json:"returns"`
	Claims                 []any      `json:"claims"`
	Refunds                []any      `json:"refunds"`
	Swaps                  []any      `json:"swaps"`
	DraftOrderId           string     `json:"draft_order_id"`
	draft_order            any        `json:"draft_order"`
	Items                  []any      `json:"items"`
	Edits                  []any      `json:"edits"`
	gift_card_transactions []any      `json:"gift_card_transactions"`
	canceled_at            *time.Time `json:"canceled_at"`
	no_notification        bool       `json:"no_notification"`
	idempotency_key        string     `json:"idempotency_key"`
	external_id            string     `json:"external_id"`
	sales_channel_id       string     `json:"sales_channel_id"`
	sales_channel          any        `json:"sales_channel"`
	shipping_total         int        `json:"shipping_total"`
	discount_total         int        `json:"discount_total"`
	tax_total              int        `json:"tax_total"`
	refunded_total         int        `json:"refunded_total"`
	total                  int        `json:"total"`
	subtotal               int        `json:"subtotal"`
	paid_total             int        `json:"paid_total"`
	refundable_amount      int        `json:"refundable_amount"`
	gift_card_total        int        `json:"gift_card_total"`
	gift_card_tax_total    int        `json:"gift_card_tax_total"`
}
