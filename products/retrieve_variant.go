package products

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type RetrieveVariantData struct {
	Variant []*schema.ProductVariant `json:"variant"`
}

type RetrieveVariantResponse struct {
	// Success response
	Data *RetrieveVariantData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

func RetrieveVariant(variantId string, config *medusa.Config) (*RetrieveVariantResponse, error) {
	path := fmt.Sprintf("/store/variants/%v", variantId)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

}
