package fiberserv

import (
	"github.com/utain/go/example/internal/adapters/presenter/fiberserv/middleware"
	"github.com/utain/go/example/internal/adapters/presenter/fiberserv/routes"
	"github.com/utain/go/example/internal/core"
	"github.com/utain/go/example/internal/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberServerOpts struct {
	Log      logs.Logging
	Services core.ServicesContainer
}

func NewFiberServer(opts FiberServerOpts) (*fiber.App, error) {
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: middleware.FiberErrorMiddleware(middleware.ErrorMiddlewareOpts{Log: opts.Log}),
	})
	fiberApp.Use(cors.New())
	fiberApp.Use(recover.New())

	apiRoute := fiberApp.Group("/api")
	routes.TodoRouter(apiRoute, opts.Services)
	return fiberApp, nil
}
