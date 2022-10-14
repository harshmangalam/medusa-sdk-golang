package customers

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
