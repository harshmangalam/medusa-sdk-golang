package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

func RetrieveVariant(variantId string, config *medusa.Config) (*schema.ProductVariant, error) {
	path := fmt.Sprintf("/store/variants/%v", variantId)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(schema.VariantResponseBody)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}

	return respBody.Variant, nil
}
