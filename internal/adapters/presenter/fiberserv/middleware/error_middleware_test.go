package middleware_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/utain/go/example/internal/adapters/presenter/fiberserv/middleware"
	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestErrorMiddleware(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.FiberErrorMiddleware(middleware.ErrorMiddlewareOpts{
			Log: logs.Nolog,
		}),
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return errs.NewError(http.StatusBadRequest, "invalid request")
	})
	app.Get("/unknown", func(c *fiber.Ctx) error {
		return errors.New("something was wrong")
	})
	app.Get("/done", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "OK!"})
	})

	t.Run("should get json error message when error is api error", func(t *testing.T) {
		res, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, `{"code": 400, "message": "invalid request"}`, string(body))
	})

	t.Run("should get json error message when error is not api error", func(t *testing.T) {
		res, err := app.Test(httptest.NewRequest(http.MethodGet, "/unknown", nil))
		assert.Nil(t, err)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, `{"code": 500, "message": "unknown error"}`, string(body))
	})

	t.Run("should not get any error message", func(t *testing.T) {
		res, err := app.Test(httptest.NewRequest(http.MethodGet, "/done", nil))
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)

		body, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		assert.JSONEq(t, `{"status": "OK!"}`, string(body))
	})
}
