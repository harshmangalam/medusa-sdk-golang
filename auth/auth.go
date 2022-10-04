package auth

import (
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

func (l *AuthSchema) Authenticate(medusa *medusa.Medusa) ([]byte, error) {
	path := "/store/auth"

	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(l).Send(medusa)

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
				medusa.SetCookie(cookie)
			}

		}
	}
	return body, nil
}

func GetSession(m *medusa.Medusa) ([]byte, error) {
	path := "/store/auth"
	resp, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(m)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DeleteSession(apiKey string) any {
	return "Ok"
}

func Exists(email string) any {
	return "exists email.."
}
