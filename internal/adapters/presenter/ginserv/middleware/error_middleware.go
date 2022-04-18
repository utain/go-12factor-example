package middleware

import (
	"net/http"

	"github.com/utain/go/example/internal/errs"
	"github.com/utain/go/example/internal/logs"

	"github.com/gin-gonic/gin"
)

const (
	tagUnhandlerError string = "[UnknownError]:"
	tagAppError       string = "[ApiError]:"
)

type ErrorOptions struct {
	Log logs.Logging
}

// GinErrorMiddleware middleware
func GinErrorMiddleware(opts ErrorOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if errors := c.Errors.ByType(gin.ErrorTypeAny); len(errors) > 0 {
			err := errors[0].Err
			if err, ok := err.(*errs.ApiError); ok {
				opts.Log.Error(err.Message, err.Data)
				c.AbortWithStatusJSON(err.Code, err)
				return
			}
			opts.Log.Error(tagUnhandlerError, logs.F{"error": err})
			c.AbortWithStatusJSON(http.StatusInternalServerError, errs.ErrUnknown)
			return
		}
	}
}
