package auth

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

// get current active user details
func GetSession(config *medusa.Config) ([]byte, error) {
	path := "/store/auth"
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

// logout current active user
func DeleteSession(config *medusa.Config) ([]byte, error) {

	path := "/store/auth"
	resp, err := request.NewRequest().SetMethod(http.MethodDelete).SetPath(path).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// check email address exists in medusa backend db
func Exists(email string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/auth/%v", email)
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
