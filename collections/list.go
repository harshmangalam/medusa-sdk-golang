package collections

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/common"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

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
func (c *CollectionsQuery) List(config *medusa.Config) ([]byte, error) {
	path := "/store/collections"

	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
