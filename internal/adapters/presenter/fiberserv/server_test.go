package fiberserv_test

import (
	"testing"

	"github.com/utain/go/example/internal/adapters/presenter/fiberserv"
	"github.com/utain/go/example/internal/core"
	"github.com/utain/go/example/internal/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiberServer(t *testing.T) {
	t.Run("new fiber server without error", func(t *testing.T) {
		engine, err := fiberserv.NewFiberServer(fiberserv.FiberServerOpts{
			Log:      logs.Nolog,
			Services: core.ServicesContainer{},
		})
		assert.Nil(t, err)
		assert.NotNil(t, engine)
		assert.IsType(t, &fiber.App{}, engine)
	})
}
