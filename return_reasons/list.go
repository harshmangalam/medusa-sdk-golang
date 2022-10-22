package returnreasons

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ListReturnReasonData struct {
	// Array of regions
	ReturnReasons []*schema.Region `json:"return_reasons"`
}

type ListReturnReasonResponse struct {
	// Success response
	Data *ListReturnReasonData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

func List(config *medusa.Config) (*ListReturnReasonResponse, error) {
	path := "/store/return-reasons"

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

}
