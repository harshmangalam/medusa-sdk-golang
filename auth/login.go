package auth

type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLogin() *LoginSchema {
	return new(LoginSchema)
}

func (l *LoginSchema) SetEmail(email string) *LoginSchema {
	l.Email = email
	return l
}

func (l *LoginSchema) SetPassword(password string) *LoginSchema {
	l.Email = password
	return l
}

func (l *LoginSchema) Send() any {
	return "login."
}
