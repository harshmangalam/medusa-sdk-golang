package products

import "github.com/harshmngalam/medusa-sdk-golang/common"

type ProductsQuery struct {
	// Query used for searching products by title, description, variant's title, variant's sku, and collection's title
	Q string `json:"q,omitempty"`

	// product IDs to search for.
	Ids []string `json:"id,omitempty"`

	// Collection IDs to search for
	CollectionIds []string `json:"collection_id"`

	// Tag IDs to search for
	Tags []string `json:"tags,omitempty"`

	// title to search for.
	Title string `json:"title,omitempty"`

	// description to search for
	Description string `json:"description,omitempty"`

	// handle to search for.
	Handle string `json:"handle,omitempty"`

	// Search for giftcards using is_giftcard=true.
	IsGiftcard bool `json:"is_giftcard,omitempty"`

	// type to search for.
	Type string `json:"type,omitempty"`

	// Date comparison for when resulting products were created.
	CreatdAt *common.DateComparison `json:"created_at,omitempty"`

	// Date comparison for when resulting products were updated.
	UpdatedAt *common.DateComparison `json:"updated_at,omitempty"`

	// How many products to skip in the result.
	Offset int `json:"offset,omitempty"`

	// Limit the number of products returned.
	Limit int `json:"limit"`

	// (Comma separated) Which fields should be expanded in each order of the result.)
	Expand string `json:"expand,omitempty"`

	// (Comma separated) Which fields should be included in each order of the result.
	Fields string `json:"fields"`
}

// create new product query
func NewProductsQuery() *ProductsQuery {
	p := new(ProductsQuery)
	p.Offset = 0
	p.Limit = 100
	return p
}

// set product query
func (p *ProductsQuery) SetQ(q string) *ProductsQuery {
	p.Q = q
	return p
}

// set product ids
func (p *ProductsQuery) SetIds(ids []string) *ProductsQuery {
	p.Ids = ids
	return p
}

// set collection ids
func (p *ProductsQuery) SetCollectionIds(collectionIds []string) *ProductsQuery {
	p.CollectionIds = collectionIds
	return p
}

// set tags
func (p *ProductsQuery) SetTags(tags []string) *ProductsQuery {
	p.Tags = tags
	return p
}

// set title
func (p *ProductsQuery) SetTitle(title string) *ProductsQuery {
	p.Title = title
	return p
}

// set description
func (p *ProductsQuery) SetDescription(description string) *ProductsQuery {
	p.Description = description
	return p
}
