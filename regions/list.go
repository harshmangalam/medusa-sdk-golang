package regions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/common"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ListRegionData struct {
	// Array of product
	Regions []*schema.Region `json:"regions"`
}

type ListRegionResponse struct {
	// Success response
	Data *ListRegionData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type RegionQuery struct {
	Offset    uint                   `json:"offset,omitempty" url:"offset,omitempty"`
	Limit     uint                   `json:"limit,omitempty" url:"limit,omitempty"`
	CreatedAt *common.DateComparison `json:"created_at,omitempty" url:"created_at,omitempty"`
	UpdatedAt *common.DateComparison `json:"updated_at,omitempty" url:"updated_at,omitempty"`
}

func NewRegionQuery() *RegionQuery {
	return new(RegionQuery)
}

func (r *RegionQuery) SetOffset(offset uint) *RegionQuery {
	r.Offset = offset
	return r
}

func (r *RegionQuery) SetLimit(limit uint) *RegionQuery {
	r.Limit = limit
	return r
}

func (r *RegionQuery) SetCreatedAt(date *common.DateComparison) *RegionQuery {
	r.CreatedAt = date
	return r
}
func (r *RegionQuery) SetUpdatedAt(date *common.DateComparison) *RegionQuery {
	r.CreatedAt = date
	return r
}

// Retrieve a list of Regions.
func (r *RegionQuery) List(config *medusa.Config) (*ListRegionResponse, error) {
	path := "/store/regions"

	qs, err := query.Values(r)
	if err != nil {
		return nil, err
	}

	parseStr := qs.Encode()

	path = fmt.Sprintf("%v?%v", path, parseStr)

	resp, err := request.
		NewRequest().
		SetMethod(http.MethodGet).
		SetPath(path).
		Send(config)

	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(ListRegionResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ListRegionData)
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
