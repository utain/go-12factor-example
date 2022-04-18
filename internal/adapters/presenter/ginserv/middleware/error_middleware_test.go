package middleware_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/utain/go/example/internal/adapters/presenter/ginserv/middleware"
	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/utain/httpheaders/methods"
)

func TestErrorMiddleware(t *testing.T) {
	app := gin.New()
	app.Use(middleware.GinErrorMiddleware(middleware.ErrorOptions{Log: logs.Nolog}))
	app.GET("/", func(c *gin.Context) {
		c.Error(errs.NewError(http.StatusBadRequest, "invalid request"))
	})
	app.GET("/unknown", func(c *gin.Context) {
		c.Error(errors.New("something was wrong"))
	})
	app.GET("/done", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK!"})
	})

	t.Run("should get json error message when error is ApiError", func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.GET, "/", nil)
		app.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, `{"code": 400, "message": "invalid request"}`, string(body))
	})

	t.Run("should get json error message when error is not ApiError", func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.GET, "/unknown", nil)
		app.ServeHTTP(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, `{"code": 500, "message": "unknown error"}`, string(body))
	})

	t.Run("should not get any error message", func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest(methods.GET, "/done", nil)
		app.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, `{"status": "OK!"}`, string(body))
	})
}
