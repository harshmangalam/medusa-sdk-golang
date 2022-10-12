package customers

import (
	"net/http"

	medusa "github.com/harshmngalam/medusa-sdk-golang"
	"github.com/harshmngalam/medusa-sdk-golang/request"
	"github.com/harshmngalam/medusa-sdk-golang/utils"
)

type CreateCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

func NewCreateCustomer() *CreateCustomer {
	return new(CreateCustomer)
}

func (c *CreateCustomer) SetFirstName(firstName string) *CreateCustomer {
	c.FirstName = firstName
	return c
}

func (c *CreateCustomer) SetLastName(lastName string) *CreateCustomer {
	c.LastName = lastName
	return c
}

func (c *CreateCustomer) SetEmail(email string) *CreateCustomer {
	c.Email = email
	return c
}

func (c *CreateCustomer) SetPassword(password string) *CreateCustomer {
	c.Password = password
	return c
}

func (c *CreateCustomer) SetPhone(phone string) *CreateCustomer {
	c.Phone = phone
	return c
}

func (c *CreateCustomer) Create(config *medusa.Config) ([]byte, error) {

	path := "/store/customers"
	resp, err := request.NewRequest().SetMethod(http.MethodPost).SetPath(path).SetData(c).Send(config)
	if err != nil {
		return nil, err
	}
	body, err := utils.ParseResponseBody(resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}
