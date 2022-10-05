package collections

import (
	"fmt"
	"net/http"
	"time"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/common"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type Collection struct {
	Title     string    `json:"title"`
	Id        string    `json:"id,omitempty"`
	Handle    string    `json:"handle,omitempty"`
	Products  []any     `json:"products,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Metadata  any       `json:"metadata,omitempty"`
}
type CollectionsQuery struct {
	Limit     int                    `json:"limit,omitempty"`
	Offset    int                    `json:"offset,omitempty"`
	CreatedAt *common.DateComparison `json:"created_at,omitempty"`
	UpdatedAt *common.DateComparison `json:"updated_at,omitempty"`
}

// create new collections query
func NewCollectionsQuery() *CollectionsQuery {
	c := new(CollectionsQuery)
	c.Limit = 10
	c.Offset = 0

	return c
}

// The number of collections to return
func (c *CollectionsQuery) SetLimit(limit int) *CollectionsQuery {
	c.Limit = limit
	return c
}

// The number of collections to skip before starting to collect the collections set
func (c *CollectionsQuery) SetOffset(offset int) *CollectionsQuery {
	c.Offset = offset
	return c
}

// Date comparison for when resulting collections were created.
func (c *CollectionsQuery) SetCreatedAt(date *common.DateComparison) *CollectionsQuery {
	c.CreatedAt = date
	return c
}

// Date comparison for when resulting collections were updated.
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

// Retrieves a Product Collection.

func Retrieve(id string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/collections/%v", id)
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
