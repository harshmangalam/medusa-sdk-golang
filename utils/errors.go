package utils

import (
	"encoding/json"

	"github.com/harshmngalam/medusa-sdk-golang/response"
)

func ParseErrors(body []byte) (*response.Errors, error) {
	respErrors := new(response.Errors)
	if err := json.Unmarshal(body, respErrors); err != nil {
		return nil, err
	}
	return respErrors, nil
}
func ParseError(body []byte) (*response.Error, error) {
	respError := new(response.Error)
	if err := json.Unmarshal(body, respError); err != nil {
		return nil, err
	}
	return respError, nil
}
func UnauthorizeError() *response.Error {
	respErr := new(response.Error)
	respErr.Message = "Unauthorized"
	return respErr
}
