package errs

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Code    int            `json:"code" example:"500"`
	Message string         `json:"message" example:"Unknown error"`
	Data    map[string]any `json:"-"`
}

func (app ApiError) Error() string {
	return app.Message
}

// Print error as JSON string
func (app ApiError) String() string {
	bs, _ := json.Marshal(app)
	return string(bs)
}

func (app *ApiError) With(name string, value any) *ApiError {
	app.Data[name] = value
	return app
}

// NewError creates a new Error instance with an optional message
func NewError(code int, message string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		Data:    map[string]any{},
	}
}

var (
	ErrTodoNotFound           = NewError(http.StatusNotFound, "todo not found")
	ErrInvalidTodoID          = NewError(http.StatusBadRequest, "invalid todo id")
	ErrInvalidTodoTitle       = NewError(http.StatusBadRequest, "invalid todo title")
	ErrInvalidTodoDescription = NewError(http.StatusBadRequest, "invalid todo description")
	ErrCannotInsertTodo       = NewError(http.StatusBadRequest, "can't add todo")
	ErrTokenInvalid           = NewError(http.StatusUnauthorized, "invalid token")
	ErrTokenExpired           = NewError(http.StatusUnauthorized, "token expired")
	ErrUnauthorized           = NewError(http.StatusUnauthorized, "please login")
	ErrUnknown                = NewError(http.StatusInternalServerError, "unknown error")
	ErrInvalidConfig          = NewError(http.StatusInternalServerError, "invalid config")
	ErrNotImplemented         = NewError(http.StatusNotImplemented, "this functional not implemented")
)
