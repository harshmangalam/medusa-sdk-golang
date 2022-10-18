package shippingoptions

type ListCartOptionsQuery struct {
	Isreturn   bool `json:"is_return,omitempty" url:"is_return,omitempty"`
	ProductIds bool `json:"product_ids,omitempty" url:"product_ids,omitempty"`
	RegionId   bool `json:"region_id,omitempty" url:"region_id,omitempty"`
}
