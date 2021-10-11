package errors

import (
	"go-example/internal/dto"
	"go-example/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrReplyUnknown = dto.ReplyError("Unknown error")

const tagUnhandlerError = "[UnhandlerError]:"
const tagAppError = "[AppError]:"

// GinError middleware
func GinError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if errors := c.Errors.ByType(gin.ErrorTypeAny); len(errors) > 0 {
			err := errors[0].Err
			if err, ok := err.(*Error); ok {
				log.Error(tagAppError, err)
				c.AbortWithStatusJSON(err.Code, err.ToReply())
				return
			}
			log.Error(tagUnhandlerError, err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrReplyUnknown)
			return
		}
	}
}
