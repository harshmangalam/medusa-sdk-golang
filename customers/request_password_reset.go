package customers

import (
	"net/http"

	medusa "github.com/harshmangalam/medusa-sdk-golang"
	"github.com/harshmangalam/medusa-sdk-golang/request"
	"github.com/harshmangalam/medusa-sdk-golang/response"
	"github.com/harshmangalam/medusa-sdk-golang/utils"
)

type RequestPasswordResetResponse struct {
	// Success response
	Data any

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type RequestPasswordReset struct {
	Email string `json:"email"`
}

func NewRequestPasswordReset() *RequestPasswordReset {
	return new(RequestPasswordReset)
}

func (r *RequestPasswordReset) SetEmail(email string) *RequestPasswordReset {
	r.Email = email
	return r
}

func (r *RequestPasswordReset) RequestReset(config *medusa.Config) (*RequestPasswordResetResponse, error) {
	path := "/store/customers/password-token"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(r).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(RequestPasswordResetResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respBody.Data = nil

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
