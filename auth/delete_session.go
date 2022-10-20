package auth

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type DeleteSessionResponse struct {
	// Success response
	Data any

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

// Logout current active user
func DeleteSession(config *medusa.Config) (*DeleteSessionResponse, error) {

	path := "/store/auth"
	resp, err := request.NewRequest().SetMethod(http.MethodDelete).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	respBody := new(DeleteSessionResponse)

	switch resp.StatusCode {
	case http.StatusOK:
		respBody.Data = nil

	case http.StatusUnauthorized:
		respErr := new(response.Error)
		respErr.Message = "Unauthorized"
		respBody.Error = respErr

	case http.StatusBadRequest:
		respErrors := new(response.Errors)
		if json.Unmarshal(body, respErrors); err != nil {
			return nil, err
		}

		if len(respErrors.Errors) == 0 {
			respError := new(response.Error)
			if json.Unmarshal(body, respError); err != nil {
				return nil, err
			}
			respBody.Error = respError
		} else {
			respBody.Errors = respErrors
		}

	default:
		respErr := new(response.Error)
		if json.Unmarshal(body, respErr); err != nil {
			return nil, err
		}
		respBody.Error = respErr

	}

	return respBody, nil
}
