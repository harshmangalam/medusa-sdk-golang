package shippingoptions

import "time"

type ShippingOption struct {
	Name         string         `json:"name"`
	RegionId     string         `json:"region_id"`
	ProfileId    string         `json:"profile_id"`
	ProviderId   string         `json:"provider_id"`
	PriceType    string         `json:"price_type"`
	Id           string         `json:"id"`
	Region       any            `json:"region"`
	Profile      any            `json:"profile"`
	Provider     any            `json:"provider"`
	Amount       int            `json:"amount"`
	IsReturn     bool           `json:"is_return"`
	Requirements []any          `json:"requirements"`
	Data         map[string]any `json:"data"`
	IncludesTax  bool           `json:"includes_tax"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    *time.Time     `json:"deleted_at"`
	Metadata     map[string]any `json:"metadata"`
}
