package api

import (
	"fmt"

	"github.com/hmoragrega/f3-payments/server/api/routing"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

const (
	f3APIDefaultIP        = ""
	f3APIDefaultPort      = 80
	f3EnvVarsPrefix       = "f3_api"
	f3APIDefaultLogFormat = "method=${method}, uri=${uri}, status=${status}\n"
)

// F3API represnts a API description
type F3API struct {
	*echo.Echo
	ip   string
	port string
}

// NewF3API buils a new API
func NewF3API() *F3API {
	initConfig()
	e := echo.New()
	setup(e)

	return &F3API{
		Echo: e,
		ip:   getIP(),
		port: getPort(),
	}
}

// Start starts the api
func (a *F3API) Start() {
	a.Echo.Logger.Fatal(a.Echo.Start(getAddress()))
}

// GetAddress Get the api address
func (a *F3API) GetAddress() string {
	return fmt.Sprintf("%s:%s", a.ip, a.port)
}

// GetEndpoint Get the api endpoint
func (a *F3API) GetEndpoint() string {
	return fmt.Sprintf("http://%s", a.GetAddress())
}

func initConfig() {
	viper.SetEnvPrefix(f3EnvVarsPrefix)
}

func setup(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: getLoggerFormat()}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	routing.RegisterRoutes(e)
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", getIP(), getPort())
}

func getIP() string {
	viper.SetDefault("ip", f3APIDefaultIP)
	viper.BindEnv("ip")

	return viper.GetString("ip")
}

func getPort() string {
	viper.SetDefault("port", f3APIDefaultPort)
	viper.BindEnv("port")

	return viper.GetString("port")
}

func getLoggerFormat() string {
	viper.SetDefault("api_log_format", f3APIDefaultLogFormat)
	viper.BindEnv("api_log_format")

	return viper.GetString("api_log_format")
}
