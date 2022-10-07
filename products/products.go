package products

import "github.com/harshmngalam/medusa-sdk-golang/common"

type ProductStatus string

const (
	Draft     ProductStatus = "draft"
	Proposed  ProductStatus = "proposed"
	Published ProductStatus = "published"
	Rejected  ProductStatus = "rejected"
)

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

type ProductVariant struct {
	Title             string           `json:"title"`
	ProductId         string           `json:"product_id"`
	InventoryQuantity int              `json:"inventory_quantity"`
	Id                string           `json:"id,omitempty"`
	Product           *Product         `json:"product"`
	Prices            []any            `json:"prices"`
	Sku               string           `json:"sku,omitempty"`
	Barcode           string           `json:"bardcode,omitempty"`
	Ean               string           `json:"ean,omitempty"`
	Upc               string           `json:"upc,omitempty"`
	VariatRank        int              `json:"variant_rank,omitempty"`
	AllowBackorder    bool             `json:"allow_backorder"`
	ManageInventory   bool             `json:"manage_inventory"`
	HsCode            string           `json:"hs_code"`
	OriginCountry     string           `json:"origin_country"`
	MidCode           string           `json:"mid_code"`
	Material          string           `json:"material"`
	Weight            int              `json:"weight"`
	Height            int              `json:"height"`
	Width             int              `json:"width"`
	Length            int              `json:"length"`
	Options           []*ProductOption `json:"options"`
	CreatedAt         string           `json:"created_at"`
	UpdatedAt         string           `json:"updated_at"`
	DeletedAt         string           `json:"deleted_at"`
	Metadata          map[string]any   `json:"metadata"`
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
	Profile       Profile           `json:"profile"`
	Weight        int               `json:"weight"`
	Height        int               `json:"height"`
	Width         int               `json:"width"`
	Length        int               `json:"length"`
	HsCode        string
	OriginCountry string         `json:"origin_country"`
	MidCode       string         `json:"mid_code"`
	Material      string         `json:"material"`
	CollectionId  string         `json:"collection_id"`
	Collection    Collection     `json:"collection"`
	TypeId        string         `json:"type_id"`
	Type          ProductType    `json:"type"`
	Tags          []Tag          `json:"tags"`
	Discountable  bool           `json:"discountable"`
	ExternalId    string         `json:"external_id"`
	SalesChannels string         `json:"sales_channels"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
	DeletedAt     string         `json:"deleted_at"`
	Metadata      map[string]any `json:"metadata"`
}
