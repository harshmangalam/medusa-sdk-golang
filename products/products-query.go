package products

import "github.com/harshmngalam/medusa-sdk-golang/common"

type ProductsQuery struct {
	Q             string                 `json:"q,omitempty"`
	Ids           []string               `json:"id,omitempty"`
	CollectionIds []string               `json:"collection_id"`
	Tags          []string               `json:"tags,omitempty"`
	Title         string                 `json:"title,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Handle        string                 `json:"handle,omitempty"`
	IsGiftcard    bool                   `json:"is_giftcard,omitempty"`
	Type          string                 `json:"type,omitempty"`
	CreatdAt      *common.DateComparison `json:"created_at,omitempty"`
	UpdatedAt     *common.DateComparison `json:"updated_at,omitempty"`
	Offset        int                    `json:"offset,omitempty"`
	Limit         int                    `json:"limit"`
	Expand        string                 `json:"expand,omitempty"`
	Fields        string                 `json:"fields"`
}

// create new product query
func NewProductsQuery() *ProductsQuery {
	p := new(ProductsQuery)
	p.Offset = 0
	p.Limit = 100
	return p
}
