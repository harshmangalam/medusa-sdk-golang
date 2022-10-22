package auth

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type AuthSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateData struct {
	Customer *schema.Customer `json:"customer,omitempty"`
}

type AuthenticateResponse struct {
	// Success response
	Data *AuthenticateData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
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

func (l *AuthSchema) Authenticate(config *medusa.Config) (*AuthenticateResponse, error) {
	path := "/store/auth"

	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(l).Send(config)

	if err != nil {
		return nil, err
	}

	body, err := utils.ParseResponseBody(resp)

	if err != nil {
		return nil, err
	}

	respBody := new(AuthenticateResponse)

	switch resp.StatusCode {
	case http.StatusOK:
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "connect.sid" {
				config.SetCookie(cookie)
			}

		}
		respData := new(AuthenticateData)
		if json.Unmarshal(body, respData); err != nil {
			return nil, err
		}
		respBody.Data = respData

	case http.StatusUnauthorized:
		respErr := utils.UnauthorizeError()
		respBody.Error = respErr
	case http.StatusBadRequest:
		respErrors, err := utils.ParseErrors(body)
		if err != nil {
			return nil, err
		}

		if len(respErrors.Errors) == 0 {

			respError, err := utils.ParseError(body)

			if err != nil {
				return nil, err
			}
			respBody.Error = respError
		} else {
			respBody.Errors = respErrors
		}

	default:
		respErr, err := utils.ParseError(body)
		if err != nil {
			return nil, err
		}
		respBody.Error = respErr

	}

	return respBody, nil
}
