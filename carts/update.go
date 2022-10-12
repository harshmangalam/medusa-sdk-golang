package carts

import "time"

type Discount struct {
	Code string `json:"code"`
}

type GiftCard struct {
	Code string `json:"code"`
}

type Address struct {
	Id          string         `json:"id"`
	CustomerId  string         `json:"customer_id"`
	Customer    []any          `json:"customer"`
	FirstName   []any          `json:"first_name"`
	LastName    []any          `json:"last_name"`
	Address1    []any          `json:"address_1"`
	Address2    []any          `json:"address_2"`
	City        []any          `json:"city"`
	CountryCode []any          `json:"country_code"`
	Country     any            `json:"country"`
	Province    string         `json:"province"`
	PostalCode  string         `json:"postal_code"`
	Phone       string         `json:"phone"`
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
	DeletedAt   time.Time      `json:"deleted_at,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}
type UpdateCart struct {
	RegionId       string         `json:"region_id"`
	CountryCode    string         `json:"country_code"`
	Email          string         `json:"email"`
	SalesChannelId string         `json:"sales_channel_id"`
	BillingAddress any            `json:"billing_address"`
	SippingAddress any            `json:"ipping_address"`
	GiftCards      []*GiftCard    `json:"gift_cards"`
	Discounts      []*Discount    `json:"discounts"`
	CustomerId     string         `json:"customer_id"`
	Context        map[string]any `json:"context"`
}
