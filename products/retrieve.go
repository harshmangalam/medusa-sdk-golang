package products

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type RetrieveQuery struct {
	// The ID of the customer's cart.
	CartId string

	// The ID of the region the customer is using. This is helpful to ensure correct prices are retrieved for a region.
	RegionId string

	// The 3 character ISO currency code to set prices based on. This is helpful to ensure correct prices are retrieved for a currency.
	CurrencyCode string
}

func Retrieve(id string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/products/%v", id)
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
