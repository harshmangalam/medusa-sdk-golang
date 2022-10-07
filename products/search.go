package products

type SearchQuery struct {
	Q      string `json:"q" url:"q"`
	Offset int    `json:"offset" url:"offset"`
	Limit  int    `json:"limit" url:"limit"`
}

func NewSearchQuery() *SearchQuery {
	return new(SearchQuery)
}

func (s *SearchQuery) SetQ(q string) *SearchQuery {
	s.Q = q
	return s
}

func (s *SearchQuery) SetOffset(offset int) *SearchQuery {
	s.Offset = offset
	return s
}

func (s *SearchQuery) SetLimit(limit int) *SearchQuery {
	s.Limit = limit
	return s
}
