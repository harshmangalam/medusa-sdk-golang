package customers

type RequestPasswordReset struct {
	Email string `json:"email"`
}

func NewRequestPasswordReset() *RequestPasswordReset {
	return new(RequestPasswordReset)
}
