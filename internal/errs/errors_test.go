package errs_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/utain/go/example/internal/errs"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	t.Run("error json tags", func(t *testing.T) {
		var err errs.ApiError

		codeTag := reflect.TypeOf(err).Field(0).Tag
		messageTag := reflect.TypeOf(err).Field(1).Tag
		errsTag := reflect.TypeOf(err).Field(2).Tag

		assert.Equal(t, "-", errsTag.Get("json"))
		assert.Equal(t, "message", messageTag.Get("json"))
		assert.Equal(t, "code", codeTag.Get("json"))
	})

	t.Run("create new error", func(t *testing.T) {
		errStr := "bad request"
		err := errs.NewError(http.StatusBadRequest, errStr)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errStr)
		assert.Empty(t, err.Data)
	})

	t.Run("create new error with error information", func(t *testing.T) {
		errStr := "bad request"
		uid := uuid.New()
		err := errs.NewError(http.StatusBadRequest, errStr).
			With("aid", uid).
			With("bid", "username")
		assert.NotNil(t, err)
		assert.EqualError(t, err, errStr)
		assert.Equal(t, map[string]any{
			"aid": uid,
			"bid": "username",
		}, err.Data)
	})
}

func TestDefinedErrors(t *testing.T) {
	t.Run("status 400: bad request", func(t *testing.T) {
		assert.Equal(t, http.StatusBadRequest, errs.ErrInvalidTodoID.Code)
		assert.Equal(t, http.StatusBadRequest, errs.ErrInvalidTodoTitle.Code)
		assert.Equal(t, http.StatusBadRequest, errs.ErrInvalidTodoDescription.Code)
		assert.Equal(t, http.StatusBadRequest, errs.ErrCannotInsertTodo.Code)
	})
	t.Run("status 401: unauthorized", func(t *testing.T) {
		assert.Equal(t, http.StatusUnauthorized, errs.ErrTokenInvalid.Code)
		assert.Equal(t, http.StatusUnauthorized, errs.ErrTokenExpired.Code)
		assert.Equal(t, http.StatusUnauthorized, errs.ErrUnauthorized.Code)
	})
	t.Run("status 404: not found", func(t *testing.T) {
		assert.Equal(t, http.StatusNotFound, errs.ErrTodoNotFound.Code)
	})
	t.Run("status 500: internal server error", func(t *testing.T) {
		assert.Equal(t, http.StatusInternalServerError, errs.ErrInvalidConfig.Code)
		assert.Equal(t, http.StatusInternalServerError, errs.ErrUnknown.Code)
	})
	t.Run("status 501: not implemented", func(t *testing.T) {
		assert.Equal(t, http.StatusNotImplemented, errs.ErrNotImplemented.Code)
	})
}
