package customers

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
