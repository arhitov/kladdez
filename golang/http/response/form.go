package response

type FormResponse struct {
	BaseResponse
	Error  *Error       `json:"error,omitempty"`
	Errors []FieldError `json:"errors,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}
