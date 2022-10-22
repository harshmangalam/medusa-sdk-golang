package schema

import "time"

type ShippingOption struct {
	Name         string         `json:"name"`
	RegionId     string         `json:"region_id"`
	ProfileId    string         `json:"profile_id"`
	ProviderId   string         `json:"provider_id"`
	PriceType    string         `json:"price_type"`
	Id           string         `json:"id,omitempty"`
	Region       any            `json:"region,omitempty"`
	Profile      any            `json:"profile,omitempty"`
	Provider     any            `json:"provider,omitempty"`
	Amount       int            `json:"amount,omitempty"`
	IsReturn     bool           `json:"is_return,omitempty"`
	Requirements []any          `json:"requirements,omitempty"`
	Data         map[string]any `json:"data,omitempty"`
	IncludesTax  bool           `json:"includes_tax,omitempty"`
	CreatedAt    *time.Time     `json:"created_at,omitempty"`
	UpdatedAt    *time.Time     `json:"updated_at,omitempty"`
	DeletedAt    *time.Time     `json:"deleted_at,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
}
