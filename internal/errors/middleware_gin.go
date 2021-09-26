package errors

import (
	"go-example/internal/dto"
	"go-example/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinError middleware
func GinError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if errors := c.Errors.ByType(gin.ErrorTypeAny); len(errors) > 0 {
			err := errors[0].Err
			if err, ok := err.(*Error); ok {
				log.Error("[AppError]:", err)
				c.AbortWithStatusJSON(err.Code, err.ToReply())
				return
			}
			log.Error("[UnhandlerError]:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ReplyError("Unknown error"))
			return
		}
	}
}
