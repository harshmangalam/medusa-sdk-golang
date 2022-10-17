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

func (l *VariantListQuery) SetIds(ids string) *VariantListQuery {
	l.Ids = ids
	return l
}

func (l *VariantListQuery) SetExpand(expand string) *VariantListQuery {
	l.Expand = expand
	return l
}

func (l *VariantListQuery) SetOffset(offset string) *VariantListQuery {
	l.Offset = offset
	return l
}

func (l *VariantListQuery) SetLimit(limit string) *VariantListQuery {
	l.Limit = limit
	return l
}

func (l *VariantListQuery) SetTitle(title string) *VariantListQuery {
	l.Title = title
	return l
}

func (l *VariantListQuery) SetInventoryQuantity(invQty any) *VariantListQuery {
	l.InventoryQuantity = invQty
	return l
}
