package errors

import (
	"time"
)

type ErroDTO struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Code      string `json:"code"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Status    uint16 `json:"status"`
	RequestID string `json:"request_id"`
}

func (e *Error) ToDTO() *ErroDTO {
	return &ErroDTO{
		ID:        e.GetID().String(),
		Timestamp: e.GetInstant().Format(time.RFC3339),
		Code:      e.GetCode(),
		Type:      e.GetTypeName(),
		Message:   e.GetMessage(),
		Status:    e.GetHTTPStatus(),
		RequestID: e.GetRequestID(),
	}
}
