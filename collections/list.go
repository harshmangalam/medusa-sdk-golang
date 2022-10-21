package collections

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/common"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ListCollectionData struct {
	Collections []*schema.Collection `json:"collections"`
	Count       uint                 `json:"count"`
	Offset      uint                 `json:"offset"`
	Limit       uint                 `json:"limit"`
}

type ListCollectionResponse struct {
	// Success response
	Data *ListCollectionData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type CollectionsQuery struct {
	// The number of collections to return
	Limit int `json:"limit,omitempty"`

	// The number of collections to skip before starting to collect the collections set
	Offset int `json:"offset,omitempty"`

	// Date comparison for when resulting collections were created.
	CreatedAt *common.DateComparison `json:"created_at,omitempty"`

	// Date comparison for when resulting collections were updated.
	UpdatedAt *common.DateComparison `json:"updated_at,omitempty"`
}

// create new collections query
func NewCollectionsQuery() *CollectionsQuery {
	c := new(CollectionsQuery)
	c.Limit = 10
	c.Offset = 0

	return c
}

// Set collection limit
func (c *CollectionsQuery) SetLimit(limit int) *CollectionsQuery {
	c.Limit = limit
	return c
}

//Set collection offset
func (c *CollectionsQuery) SetOffset(offset int) *CollectionsQuery {
	c.Offset = offset
	return c
}

// Set collection created date
func (c *CollectionsQuery) SetCreatedAt(date *common.DateComparison) *CollectionsQuery {
	c.CreatedAt = date
	return c
}

// Set collection updated date
func (c *CollectionsQuery) SetUpdatedAt(date *common.DateComparison) *CollectionsQuery {
	c.UpdatedAt = date
	return c
}

// Retrieve a list of Product Collection.
func (c *CollectionsQuery) List(config *medusa.Config) (*ListCollectionResponse, error) {
	path := "/store/collections"

	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	respBody := new(ListCollectionResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ListCollectionData)
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
