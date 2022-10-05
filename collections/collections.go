package collections

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/common"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type CollectionsQuery struct {
	Limit     int                            `json:"limit,omitempty"`
	Offset    int                            `json:"offset,omitempty"`
	CreatedAt *common.DateComparisonOperator `json:"created_at,omitempty"`
	UpdatedAt *common.DateComparisonOperator `json:"updated_at,omitempty"`
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

// Retrieve a list of Product Collection.

func NewCollectionsQuery() *CollectionsQuery {
	c := new(CollectionsQuery)
	c.Limit = 10
	c.Offset = 0

	return c
}

func (c *CollectionsQuery) SetLimit(limit int) *CollectionsQuery {
	c.Limit = limit
	return c
}

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
