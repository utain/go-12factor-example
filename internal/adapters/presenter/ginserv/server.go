package ginserv

import (
	"github.com/utain/go/example/internal/adapters/presenter/ginserv/middleware"
	"github.com/utain/go/example/internal/adapters/presenter/ginserv/routes"
	"github.com/utain/go/example/internal/core"
	"github.com/utain/go/example/internal/logs"

	"github.com/gin-gonic/gin"
)

type GinServerOpts struct {
	Log      logs.Logging
	Services core.ServicesContainer
}

func NewGinServer(opts GinServerOpts) (*gin.Engine, error) {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(middleware.GinErrorMiddleware(middleware.ErrorOptions{
		Log: opts.Log,
	}))

	api := router.Group("/api")
	routes.TodoRouter(api, opts.Services.TodoServicePort)
	return router, nil
}
