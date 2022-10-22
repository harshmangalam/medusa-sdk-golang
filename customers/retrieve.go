package customers

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type RetrieveCustomerData struct {
	Customer *schema.Customer `json:"customer"`
}

type RetrieveCustomerResponse struct {
	// Success response
	Data *RetrieveCustomerData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

func Retrieve(config *medusa.Config) (*RetrieveCustomerResponse, error) {
	path := "/store/customers/me"
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

}
