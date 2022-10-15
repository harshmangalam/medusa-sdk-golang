package orders

type Lookup struct {
	display_id       string  `json:"display_id" url:"display_id"`
	email            string  `json:"email" url:"email"`
	shipping_address Address `json:"shipping_address"`
}
