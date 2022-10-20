//Get Current Customer

package auth

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/customers"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type GetSessionData struct {
	Customer customers.Customer `json:"customer,omitempty"`
}

type GetSessionErrors struct {
	// Array of errors
	Errors []*response.Error `json:"errors,omitempty"`
	// Default: "Provided request body contains errors. Please check the data and retry the request"
	Message string `json:"message,omitempty"`
}

type GetSessionResponse struct {
	// Success response
	Data *GetSessionData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *GetSessionErrors
}

// Gets the currently logged in Customer.
func GetSession(config *medusa.Config) (*GetSessionResponse, error) {
	path := "/store/auth/"
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(GetSessionResponse)

	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(GetSessionData)
		if json.Unmarshal(body, respData); err != nil {
			return nil, err
		}
		respBody.Data = respData

	case http.StatusUnauthorized:
		respErr := new(response.Error)
		respErr.Message = "Unauthorized"
		respBody.Error = respErr

	case http.StatusBadRequest:
		respErrors := new(GetSessionErrors)
		if json.Unmarshal(body, respErrors); err != nil {
			return nil, err
		}
		respBody.Errors = respErrors

	default:
		respErr := new(response.Error)
		if json.Unmarshal(body, respErr); err != nil {
			return nil, err
		}
		respBody.Error = respErr

	}

	return respBody, err
}
