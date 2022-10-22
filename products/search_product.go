package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type SearchProductData struct {
	// Array of results. The format of the items depends on the search engine installed on the server.
	Hits []any `json:"hits"`
}

type SearchProductResponse struct {
	// Success response
	Data *SearchProductData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type SearchQuery struct {
	// The query to run the search with.
	Q string `json:"q,omitempty" url:"q,omitempty"`

	// How many products to skip in the result.
	Offset int `json:"offset,omitempty" url:"offset,omitempty"`

	// Limit the number of products returned.
	Limit int `json:"limit,omitempty" url:"limit,omitempty"`
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

// Run a search query on products using the search engine installed on Medusa
func (s *SearchQuery) Search(config *medusa.Config) (*SearchProductResponse, error) {
	path := "/store/products/search"

	qs, err := query.Values(s)
	if err != nil {
		return nil, err
	}

	parseStr := qs.Encode()

	path = fmt.Sprintf("%v?%v", path, parseStr)
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
	respBody := new(SearchProductResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(SearchProductData)
		if err := json.Unmarshal(body, respData); err != nil {
			return nil, err
		}
		respBody.Data = respData

	case http.StatusUnauthorized:
		respErr := utils.UnauthorizeError()
		respBody.Error = respErr

	case http.StatusBadRequest:
		respErrors, err := utils.ParseErrors(body)
		if err != nil {
			return nil, err
		}
		if len(respErrors.Errors) == 0 {
			respError, err := utils.ParseError(body)
			if err != nil {
				return nil, err
			}
			respBody.Error = respError
		} else {
			respBody.Errors = respErrors
		}

	default:
		respErr, err := utils.ParseError(body)
		if err != nil {
			return nil, err
		}
		respBody.Error = respErr
	}

	return respBody, nil
}
