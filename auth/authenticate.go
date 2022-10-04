package auth

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
	l.Email = password
	return l
}

func (l *AuthSchema) Authenticate() any {
	return "authenticate..."
}
