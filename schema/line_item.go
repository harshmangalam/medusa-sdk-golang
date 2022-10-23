package schema

import "time"

type LineItem struct {
	Title             string          `json:"title,omitempty"`
	UnitPrice         bool            `json:"unit_price,omitempty"`
	Quantity          uint            `json:"quantity,omitempty"`
	Id                string          `json:"id,omitempty"`
	CartId            string          `json:"cart_id,omitempty"`
	Cart              *Cart           `json:"cart,omitempty"`
	OrderId           string          `json:"order_id,omitempty"`
	Order             *Order          `json:"order,omitempty"`
	SwapId            string          `json:"swap_id,omitempty"`
	Swap              *Swap           `json:"swap,omitempty"`
	ClaimOrderId      string          `json:"claim_order_id,omitempty"`
	ClaimOrder        any             `json:"claim_order,omitempty"`
	TaxLines          []any           `json:"tax_lines,omitempty"`
	Adjustments       []any           `json:"adjustments,omitempty"`
	Description       string          `json:"description,omitempty"`
	Thumbnail         string          `json:"thumbnail,omitempty"`
	IsReturn          bool            `json:"is_return,omitempty"`
	IsGiftcard        bool            `json:"is_giftcard,omitempty"`
	ShouldMerge       bool            `json:"should_merge,omitempty"`
	AllowDiscounts    bool            `json:"allow_discounts,omitempty"`
	HasShipping       bool            `json:"has_shipping,omitempty"`
	VariantId         string          `json:"variant_id,omitempty"`
	Variant           *ProductVariant `json:"variant,omitempty"`
	FulfilledQuantity uint            `json:"fulfilled_quantity,omitempty"`
	ReturnedQuantity  uint            `json:"returned_quantity,omitempty"`
	Shipped_quantity  uint            `json:"shipped_quantity,omitempty"`
	Refundable        uint            `json:"refundable,omitempty"`
	Subtotal          uint            `json:"subtotal,omitempty"`
	TaxTotal          uint            `json:"tax_total,omitempty"`
	Total             uint            `json:"total,omitempty"`
	OriginalTotal     uint            `json:"original_total,omitempty"`
	OriginalTaxTotal  uint            `json:"original_tax_total,omitempty"`
	DiscountTotal     uint            `json:"discount_total,omitempty"`
	GiftCardTotal     uint            `json:"gift_card_total,omitempty"`
	CreatedAt         *time.Time      `json:"created_at,omitempty"`
	DeletedAt         *time.Time      `json:"deleted_at,omitempty"`
	Metadata          map[string]any  `json:"metadata,omitempty"`
}
