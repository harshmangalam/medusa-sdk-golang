package products

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type SearchQuery struct {
	Q      string `json:"q,omitempty" url:"q,omitempty"`
	Offset int    `json:"offset,omitempty" url:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty" url:"limit,omitempty"`
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

func (s *SearchQuery) Search(config *medusa.Config) ([]byte, error) {
	path := "/store/products/search"

	qs, err := query.Values(s)
	if err != nil {
		return nil, err
	}

	parseStr := qs.Encode()

	path = fmt.Sprintf("%v?%v", path, parseStr)

	fmt.Println(path)
	resp, err := request.
		NewRequest().
		SetMethod(http.MethodPost).
		SetPath(path).
		Send(config)

	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
