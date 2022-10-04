package auth

import (
	"fmt"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type AuthSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuth() *AuthSchema {
	return new(AuthSchema)
}

func (l *AuthSchema) SetEmail(email string) *AuthSchema {
	l.Email = email
	return l
}

func (l *AuthSchema) SetPassword(password string) *AuthSchema {
	l.Password = password
	return l
}

func (l *AuthSchema) Authenticate(config *medusa.Config) ([]byte, error) {
	path := "/store/auth"

	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(l).Send(config)

	if err != nil {
		return nil, err
	}

	body, err := utils.ParseResponseBody(resp)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {

		for _, cookie := range resp.Cookies() {
			if cookie.Name == "connect.sid" {
				config.SetCookie(cookie)
			}

		}
	}
	return body, nil
}

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

func Exists(email string, config *medusa.Config) ([]byte, error) {
	path := fmt.Sprintf("/store/auth/%v", email)
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
