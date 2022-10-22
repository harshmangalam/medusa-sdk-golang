package schema

import (
	"time"

	"github.com/harshmngalam/medusa-sdk-golang/common"
)

type ProductStatus string

const (
	Draft     ProductStatus = "draft"
	Proposed  ProductStatus = "proposed"
	Published ProductStatus = "published"
	Rejected  ProductStatus = "rejected"
)

type ShippingOptions struct {
	CreatedAt string         `json:"created_at,omitempty"`
	UpdatedAt string         `json:"updated_at,omitempty"`
	DeletedAt string         `json:"deleted_at,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}
type ShippingProfile struct {
	Name            string         `json:"name"`
	Type            string         `json:"type"`
	Id              string         `json:"id"`
	Products        []*Product     `json:"products"`
	ShippingOptions any            `json:"shipping_options"`
	CreatedAt       string         `json:"created_at,omitempty"`
	UpdatedAt       string         `json:"updated_at,omitempty"`
	DeletedAt       string         `json:"deleted_at,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
}
type ProductOption struct {
	Title     string         `json:"title"`
	ProductId string         `json:"product_id"`
	Id        string         `json:"id,omitempty"`
	Values    []any          `json:"values"`
	Product   *Product       `json:"product"`
	CreatedAt string         `json:"created_at,omitempty"`
	UpdatedAt string         `json:"updated_at,omitempty"`
	DeletedAt string         `json:"deleted_at,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

type ProductType struct {
	Value     string         `json:"value"`
	Id        string         `json:"id"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt string         `json:"deleted_at"`
	Metadata  map[string]any `json:"metadata"`
}

type Tag struct {
	Value     string         `json:"value"`
	Id        string         `json:"id"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt string         `json:"deleted_at"`
	Metadata  map[string]any `json:"metadata"`
}

type Product struct {
	Title         string            `json:"title"`
	ProfileId     string            `json:"profile_id"`
	Id            string            `json:"id"`
	Subtitle      string            `json:"subtitle"`
	Description   string            `json:"description"`
	Handle        string            `json:"handle"`
	IsGiftcard    bool              `json:"is_giftcard"`
	Status        ProductStatus     `json:"status"`
	Images        []*common.Image   `json:"images"`
	Thumbnail     string            `json:"thumbnail"`
	Options       []*ProductOption  `json:"options"`
	Variants      []*ProductVariant `json:"variants"`
	Profile       *ShippingProfile  `json:"profile"`
	Weight        int               `json:"weight"`
	Height        int               `json:"height"`
	Width         int               `json:"width"`
	Length        int               `json:"length"`
	HsCode        string
	OriginCountry string         `json:"origin_country"`
	MidCode       string         `json:"mid_code"`
	Material      string         `json:"material"`
	CollectionId  string         `json:"collection_id"`
	Collection    *Collection    `json:"collection"`
	TypeId        string         `json:"type_id"`
	Type          *ProductType   `json:"type"`
	Tags          []*Tag         `json:"tags"`
	Discountable  bool           `json:"discountable"`
	ExternalId    string         `json:"external_id"`
	SalesChannels string         `json:"sales_channels"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     *time.Time     `json:"deleted_at"`
	Metadata      map[string]any `json:"metadata"`
}
