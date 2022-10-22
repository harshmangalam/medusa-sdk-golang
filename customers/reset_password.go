package customers

import (
	"encoding/json"
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/response"
	"github.com/harshmngalam/medusa-sdk-golang/schema"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type ResetPasswordData struct {
	Customer *schema.Customer `json:"customer"`
}

type ResetPasswordResponse struct {
	// Success response
	Data *ResetPasswordData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type ResetPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func NewResetPassword() *ResetPassword {
	return new(ResetPassword)
}

func (r *ResetPassword) SetEmail(email string) *ResetPassword {
	r.Email = email
	return r
}

func (r *ResetPassword) SetPassword(password string) *ResetPassword {
	r.Password = password
	return r
}

func (r *ResetPassword) SetToken(token string) *ResetPassword {
	r.Token = token
	return r
}

func (r *ResetPassword) Reset(config *medusa.Config) (*ResetPasswordResponse, error) {
	path := "/store/customers/password-reset"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(r).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(ResetPasswordResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(ResetPasswordData)
		if err := json.Unmarshal(body, respData); err != nil {
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
