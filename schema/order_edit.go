package schema

import "time"

type OrderEdit struct {
	// The ID of the order that is edited
	OrderId string `json:"order_id"`

	// Represents an order
	Order *Order `json:"order"`

	// Line item changes array.
	Changes []any `json:"changes"`

	// The unique identifier of the user or customer who created the order edit.
	CreatedBy string `json:"created_by"`

	// The order edit's ID
	Id string `json:"id,omitempty"`

	// An optional note with additional details about the order edit.
	InternalNote string `json:"internal_note,omitempty"`

	// The unique identifier of the user or customer who requested the order edit.
	RequestedBy string `json:"requested_by,omitempty"`

	// The date with timezone at which the edit was requested.
	RequestedAt *time.Time `json:"requested_at,omitempty"`

	// The unique identifier of the user or customer who confirmed the order edit
	ConfirmedBy string `json:"confirmed_by,omitempty"`

	// The date with timezone at which the edit was confirmed.
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`

	// The date with timezone at which the edit was declined.
	DeclinedBy string `json:"declined_by,omitempty"`

	// An optional note why the order edit is declined.
	DeclinedAt *time.Time `json:"declined_at,omitempty"`

	// An optional note why the order edit is declined.
	DeclinedReason string `json:"declined_reason,omitempty"`

	// The total of subtotal
	Subtotal uint `json:"subtotal,omitempty"`

	// The total of discount
	DiscountTotal uint `json:"discount_total,omitempty"`

	// The total of the shipping amount
	ShippingTotal uint `json:"ShippingTotal,omitempty"`

	// The total of the gift card amount
	GiftCardTotal uint `json:"gift_card_total,omitempty"`

	// The total of the gift card tax amount
	GiftCardTaxTotal uint `json:"gift_card_tax_total,omitempty"`

	// The total of tax
	TaxTotal uint `json:"tax_total,omitempty"`

	// The total amount of the edited order.
	Total uint `json:"total,omitempty"`

	// The difference between the total amount of the order and total amount of edited order.
	DifferenceDue uint `json:"difference_due,omitempty"`

	// Computed line items from the changes.
	Items []any `json:"items,omitempty"`
}
