package ginserv_test

import (
	"testing"

	"github.com/utain/go/example/internal/adapters/presenter/ginserv"
	"github.com/utain/go/example/internal/core"
	"github.com/utain/go/example/internal/logs"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGinServer(t *testing.T) {
	t.Run("new gin server without error", func(t *testing.T) {
		engine, err := ginserv.NewGinServer(ginserv.GinServerOpts{
			Log:      logs.Nolog,
			Services: core.ServicesContainer{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, engine)
		assert.IsType(t, &gin.Engine{}, engine)
	})
}
