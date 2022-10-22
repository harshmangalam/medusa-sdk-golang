package schema

import (
	"time"
)

type VariantResponseBody struct {
	Variant *ProductVariant `json:"variant"`
}
type ProductVariant struct {
	Title                  string         `json:"title"`
	ProductId              string         `json:"product_id"`
	InventoryQuantity      string         `json:"inventory_quantity"`
	Id                     string         `json:"id,omitempty"`
	Product                *Product       `json:"product,omitempty"`
	Prices                 []any          `json:"prices,omitempty"`
	Sku                    string         `json:"sku,omitempty"`
	Barcode                string         `json:"barcode,omitempty"`
	Ean                    string         `json:"ean,omitempty"`
	Upc                    string         `json:"upc,omitempty"`
	VariantRank            int            `json:"variant_rank,omitempty"`
	AllowBackorder         bool           `json:"allow_backorder,omitempty"`
	ManageInventory        bool           `json:"manage_inventory,omitempty"`
	HsCode                 string         `json:"hs_code,omitempty"`
	MidCode                string         `json:"mid_code,omitempty"`
	Material               string         `json:"material,omitempty"`
	Weight                 int            `json:"weight,omitempty"`
	Height                 int            `json:"height,omitempty"`
	Width                  int            `json:"width,omitempty"`
	Length                 int            `json:"length,omitempty"`
	Options                []any          `json:"options,omitempty"`
	CreatedAt              *time.Time     `json:"created_at,omitempty"`
	UpdatedAt              *time.Time     `json:"updated_at,omitempty"`
	DeletedAt              *time.Time     `json:"deleted_at,omitempty"`
	Metadata               map[string]any `json:"metadata,omitempty"`
	OriginalPrice          int            `json:"original_price,omitempty"`
	CalculatedPrice        int            `json:"calculated_price,omitempty"`
	OriginalPriceInclTax   int            `json:"original_price_incl_tax,omitempty"`
	CalculatedPriceInclTax int            `json:"calculated_price_incl_tax,omitempty"`
	OriginalTax            int            `json:"original_tax,omitempty"`
	CalculatedTax          int            `json:"calculated_tax,omitempty"`
	TaxRates               []any          `json:"tax_rates,omitempty"`
}
