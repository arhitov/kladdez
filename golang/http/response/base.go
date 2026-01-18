package response

import "time"

type StatusResponse string

const (
	StatusSuccess StatusResponse = "success"
	StatusError   StatusResponse = "error"
)

// BaseResponse Базовый ответ с оберткой
type BaseResponse struct {
	Status    StatusResponse `json:"status"`
	Timestamp time.Time      `json:"timestamp"`
	Message   string         `json:"message,omitempty"`
}
