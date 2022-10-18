package shippingoptions

type ListCartOptionsQuery struct {
	Isreturn   bool   `json:"is_return,omitempty" url:"is_return,omitempty"`
	ProductIds string `json:"product_ids,omitempty" url:"product_ids,omitempty"`
	RegionId   string `json:"region_id,omitempty" url:"region_id,omitempty"`
}

func NewListCartOptionsQuery() *ListCartOptionsQuery {
	return new(ListCartOptionsQuery)
}

func (l *ListCartOptionsQuery) SetIsreturn(isReturn bool) *ListCartOptionsQuery {
	l.Isreturn = isReturn
	return l
}

func (l *ListCartOptionsQuery) SetProductIds(productIds string) *ListCartOptionsQuery {
	l.ProductIds = productIds
	return l
}

func (l *ListCartOptionsQuery) SetRegionId(regionId string) *ListCartOptionsQuery {
	l.RegionId = regionId
	return l
}
