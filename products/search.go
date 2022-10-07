package products

type SearchQuery struct {
	Q      string `json:"q" url:"q"`
	Offset int    `json:"offset" url:"offset"`
	Limit  int    `json:"limit" url:"limit"`
}
