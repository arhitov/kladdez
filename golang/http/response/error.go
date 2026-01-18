package response

type ErrorResponse struct {
	BaseResponse
	Error Error `json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}
