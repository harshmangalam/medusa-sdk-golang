package regions

import "time"

type Region struct {
	Name                 string         `json:"name"`
	CurrencyCode         string         `json:"currency_code"`
	TaxRate              int            `json:"tax_rate"`
	Id                   string         `json:"id"`
	Currency             any            `json:"currency"`
	TaxRates             []any          `json:"tax_rates"`
	TaxCode              string         `json:"tax_code"`
	GiftCardsTaxable     bool           `json:"gift_cards_taxable"`
	AutomaticTaxes       bool           `json:"automatic_taxes"`
	Countries            []any          `json:"countries"`
	TaxProviderId        string         `json:"tax_provider_id"`
	TaxProvider          any            `json:"tax_provider"`
	PaymentProviders     []any          `json:"payment_providers"`
	FulfillmentProviders []any          `json:"fulfillment_providers"`
	IncludesTax          string         `json:"includes_tax"`
	CreatedAt            *time.Time     `json:"created_at"`
	UpdatedAt            *time.Time     `json:"updated_at"`
	DeletedAt            *time.Time     `json:"deleted_at"`
	Metadata             map[string]any `json:"metadata"`
}

type RetrieveResponseBody struct {
	Region *Region `json:"region"`
}
