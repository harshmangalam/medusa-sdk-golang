package response

type Error struct {
	// A slug indicating the type of the error
	Type string `json:"type"`

	// Description of the error that occurred.
	Message string `json:"message"`

	// A slug code to indicate the type of the error.
	Code string `json:"code,omitempty"`
}

type Errors struct {
	// Array of errors
	Errors []*Error `json:"errors,omitempty"`
	// Default: "Provided request body contains errors. Please check the data and retry the request"
	Message string `json:"message,omitempty"`
}
