package response

type Error struct {
	// A slug indicating the type of the error
	Type string `json:"type"`

	// Description of the error that occurred.
	Message string `json:"message"`

	// A slug code to indicate the type of the error.
	Code string `json:"code,omitempty"`
}
