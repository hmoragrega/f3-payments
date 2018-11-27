package main

import (
	"context"

	"github.com/hmoragrega/f3-payments/cmd/api/config"
	"github.com/hmoragrega/f3-payments/cmd/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// F3API represnts a API description
type F3API struct {
	c *config.Config
	*echo.Echo
}

// NewF3API buils a new API
func NewF3API(c *config.Config, d *config.DIC) *F3API {
	f3 := &F3API{c, echo.New()}
	f3.setup()
	config.RegisterRoutes(f3.Echo, d)
	f3.Echo.HTTPErrorHandler = handlers.JSONApiErrorPrettyHanler

	return f3
}

// Start starts the api
func (a *F3API) Start() {
	a.Echo.Start(a.c.GetServerAddress())
}

// Shutdown terminates the api in a clean way
func (a *F3API) Shutdown(ctx context.Context) error {
	return a.Echo.Shutdown(ctx)
}

func (a *F3API) setup() {
	a.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: a.c.LogFormat}))
	a.Echo.Use(middleware.Recover())
	a.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}
