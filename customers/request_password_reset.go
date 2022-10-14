package customers

type RequestPasswordReset struct {
	Email string `json:"email"`
}
