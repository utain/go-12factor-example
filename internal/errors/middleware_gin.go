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
		errors := c.Errors.ByType(gin.ErrorTypeAny)
		if len(errors) > 0 {
			err := errors[0].Err
			if err, ok := err.(*Error); ok {
				log.Error("[AppError]:", err)
				c.AbortWithStatusJSON(err.Code, dto.ReplyError(err.Message))
				return
			}
			log.Error("[AppErrorUnhandle]:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ReplyError("Unknown error"))
			return
		}
	}
}
