package middleware

import (
	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/gofiber/fiber/v2"
)

type ErrorMiddlewareOpts struct {
	Log logs.Logging
}

func FiberErrorMiddleware(opts ErrorMiddlewareOpts) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if err, ok := err.(*errs.ApiError); ok {
			opts.Log.Error(err.Message, err.Data)
			return c.Status(err.Code).JSON(err)
		}
		opts.Log.Error("Have unknown error", logs.F{"error": err})
		return c.Status(errs.ErrUnknown.Code).JSON(errs.ErrUnknown)
	}
}
