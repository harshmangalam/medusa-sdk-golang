package regions

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type RetrieveRegiondata struct {
	Region []*schema.Region `json:"region"`
}

type RetrieveRegionResponse struct {
	// Success response
	Data *RetrieveRegiondata

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

func Retrieve(regionId string, config *medusa.Config) (*RetrieveRegionResponse, error) {
	path := fmt.Sprintf("/store/regions/%v", regionId)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

}
