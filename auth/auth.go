package auth

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
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

	res, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(l).Send(medusa)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetSession(m *medusa.Medusa) ([]byte, error) {
	path := "/store/auth"
	res, err := request.NewRequest().SetMethod(http.MethodGet).SetPath(path).Send(m)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func DeleteSession(apiKey string) any {
	return "Ok"
}

func Exists(email string) any {
	return "exists email.."
}
