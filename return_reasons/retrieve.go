package returnreasons

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type RetrieveReturnReasonData struct {
	ReturnReason []*schema.ReturnReason `json:"return_reason"`
}

type RetrieveReturnReasonResponse struct {
	// Success response
	Data *RetrieveReturnReasonData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

func Retrieve(id string, config *medusa.Config) (*RetrieveReturnReasonResponse, error) {
	path := fmt.Sprintf("/store/return-reasons/%v", id)
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

}
