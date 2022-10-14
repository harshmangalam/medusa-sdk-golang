package customers

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
