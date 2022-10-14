package customers

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

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

func (r *ResetPassword) Send(config *medusa.Config) ([]byte, error) {
	path := "/store/customers/password-token"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(r).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	return body, nil
}
