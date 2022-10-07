package products

type ProductStatus string

const (
	Draft     ProductStatus = "draft"
	Proposed  ProductStatus = "proposed"
	Published ProductStatus = "published"
	Rejected  ProductStatus = "rejected"
)

type Product struct {
	Title         string        `json:"title"`
	ProfileId     string        `json:"profile_id"`
	Id            string        `json:"id"`
	Subtitle      string        `json:"subtitle"`
	Description   string        `json:"description"`
	Handle        string        `json:"handle"`
	IsGiftcard    bool          `json:"is_giftcard"`
	Status        ProductStatus `json:"status"`
	Images        []Image       `json:"images"`
	Thumbnail     string        `json:"thumbnail"`
	Options       []Option      `json:"options"`
	Variants      []Variant     `json:"variants"`
	Profile       Profile       `json:"profile"`
	Weight        int           `json:"weight"`
	Height        int           `json:"height"`
	Width         int           `json:"width"`
	Length        int           `json:"length"`
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
