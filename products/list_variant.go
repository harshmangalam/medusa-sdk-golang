package products

type VariantListQuery struct {
	Ids               string `json:"ids" url:"ids"`
	Expand            string `json:"expand" url:"expand"`
	Offset            string `json:"offset" url:"offset"`
	Limit             string `json:"limit" url:"limit"`
	Title             string `json:"title" url:"title"`
	InventoryQuantity any    `json:"inventory_quantity" url:"inventory_quantity"`
}

func NewVariantListQuery() *VariantListQuery {
	return new(VariantListQuery)
}

func (l *VariantListQuery) SetIds() *VariantListQuery {

	return l
}

func (l *VariantListQuery) SetExpand() *VariantListQuery {

	return l
}

func (l *VariantListQuery) SetOffset() *VariantListQuery {

	return l
}

func (l *VariantListQuery) SetLimit() *VariantListQuery {

	return l
}

func (l *VariantListQuery) SetTitle() *VariantListQuery {

	return l
}

func (l *VariantListQuery) SetInventoryQuantity(invQty any) *VariantListQuery {
	l.InventoryQuantity = invQty
	return l
}
