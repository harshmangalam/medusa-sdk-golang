package auth

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/customers"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type AuthenticateResponse struct {
	Customer *customers.Customer `json:"customer"`
}

type AuthSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// create new auth struct and return thier reference for further method chaining
func NewAuth() *AuthSchema {
	return new(AuthSchema)
}

// set customer email address
func (l *AuthSchema) SetEmail(email string) *AuthSchema {
	l.Email = email
	return l
}

// set customer password
func (l *AuthSchema) SetPassword(password string) *AuthSchema {
	l.Password = password
	return l
}

// make api post request to medusa api and authenticate customer using provided email and password
func (l *AuthSchema) Authenticate(config *medusa.Config) (*customers.Customer, error) {
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

	respBody := new(AuthenticateResponse)

	if err := json.Unmarshal(body, respBody); err != nil {
		return nil, err
	}
	return respBody.Customer, nil
}
