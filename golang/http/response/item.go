package response

type ItemResponse[T any] struct {
	BaseResponse
	Data  T      `json:"item"`
	Error *Error `json:"error,omitempty"`
}
