package orderedits

import (
	"time"

	"github.com/harshmngalam/medusa-sdk-golang/orders"
)

type OrderEdit struct {
	OrderId          string        `json:"order_id"`
	Order            *orders.Order `json:"order"`
	Changes          []any         `json:"changes"`
	CreatedBy        string        `json:"created_by"`
	Id               string        `json:"id,omitempty"`
	InternalNote     string        `json:"internal_note,omitempty"`
	RequestedBy      string        `json:"requested_by,omitempty"`
	ConfirmedBy      string        `json:"confirmed_by,omitempty"`
	ConfirmedAt      *time.Time    `json:"confirmed_at,omitempty"`
	DeclinedBy       string        `json:"declined_by,omitempty"`
	DeclinedAt       *time.Time    `json:"declined_at,omitempty"`
	DeclinedReason   string        `json:"declined_reason,omitempty"`
	Subtotal         int           `json:"subtotal,omitempty"`
	DiscountTotal    int           `json:"discount_total,omitempty"`
	ShippingTotal    int           `json:"ShippingTotal,omitempty"`
	GiftCardTotal    int           `json:"gift_card_total,omitempty"`
	GiftCardTaxTotal int           `json:"gift_card_tax_total,omitempty"`
	TaxTotal         int           `json:"tax_total,omitempty"`
	Total            int           `json:"total,omitempty"`
	Items            []any         `json:"items,omitempty"`
}
