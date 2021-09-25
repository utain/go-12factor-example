package errors

import (
	"go-example/internal/dto"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (app *Error) Error() string {
	return app.Message
}

func (app *Error) ToReply() *dto.ErrorReply {
	return &dto.ErrorReply{
		Error: dto.ErrorMessage{
			Message: app.Message,
		},
	}
}

// NewError creates a new Error instance with an optional message
func NewError(code int, message string) *Error {
	e := &Error{
		Code:    code,
		Message: message,
	}
	return e
}
