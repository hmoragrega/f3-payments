package main

import (
	"github.com/hmoragrega/f3-payments/cmd/api/config"
	"github.com/hmoragrega/f3-payments/cmd/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// F3API represnts a API description
type F3API struct {
	c *config.Config
	d *config.DIC
	*echo.Echo
}

// NewF3API buils a new API
func NewF3API(c *config.Config, d *config.DIC) *F3API {
	f3 := &F3API{c, d, echo.New()}
	f3.setup()
	f3.Echo.HTTPErrorHandler = handlers.JSONApiErrorPrettyHanler

	return f3
}

// Start starts the api
func (a *F3API) Start() {
	run := func() error {
		if err := config.RegisterRoutes(a.Echo, a.d); err != nil {
			return err
		}

		return a.Echo.Start(a.c.GetServerAddress())
	}

	a.Echo.Logger.Fatal(run())
}

func (a *F3API) setup() {
	a.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: a.c.LogFormat}))
	a.Echo.Use(middleware.Recover())
	a.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}
