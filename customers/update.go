package customers

type UpdateCustomer struct {
	Email          string         `json:"email"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	BillingAddress any            `json:"billing_address"`
	Phone          string         `json:"phone"`
	Password       string         `json:"password"`
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
