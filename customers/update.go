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

type UpdateCustomerData struct {
	Customer *schema.Customer `json:"customer"`
}

type UpdateCustomerResponse struct {
	// Success response
	Data *UpdateCustomerData

	// Error response
	Error *response.Error

	// Errors in case of multiple errors
	Errors *response.Errors
}

type UpdateCustomer struct {
	Email          string         `json:"email,omitempty"`
	FirstName      string         `json:"first_name,omitempty"`
	LastName       string         `json:"last_name,omitempty"`
	BillingAddress any            `json:"billing_address,omitempty"`
	Phone          string         `json:"phone,omitempty"`
	Password       string         `json:"password,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
}

func NewUpdateCustomer() *UpdateCustomer {

	return new(UpdateCustomer)
}
func (u *UpdateCustomer) SetEmail(email string) *UpdateCustomer {
	u.Email = email
	return u
}
func (u *UpdateCustomer) SetFirstName(firstName string) *UpdateCustomer {
	u.FirstName = firstName
	return u
}

func (u *UpdateCustomer) SetLastName(lastName string) *UpdateCustomer {
	u.LastName = lastName
	return u
}

func (u *UpdateCustomer) SetBillingAddress(billingAddr any) *UpdateCustomer {
	u.BillingAddress = billingAddr
	return u
}

func (u *UpdateCustomer) SetPhone(phone string) *UpdateCustomer {
	u.Phone = phone
	return u
}

func (u *UpdateCustomer) SetPassword(password string) *UpdateCustomer {
	u.Password = password
	return u
}

func (u *UpdateCustomer) SetMetadata(metaData map[string]any) *UpdateCustomer {
	u.Metadata = metaData
	return u
}

func (u *UpdateCustomer) Update(config *medusa.Config) (*UpdateCustomerResponse, error) {
	path := "/store/customers/me"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(u).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	respBody := new(UpdateCustomerResponse)
	switch resp.StatusCode {
	case http.StatusOK:
		respData := new(UpdateCustomerData)
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
