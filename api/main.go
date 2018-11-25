package main

import (
	"fmt"

	"github.com/hmoragrega/f3-payments/api/routing"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

// F3APIDefaultIP Default IP for the api server
const F3APIDefaultIP = ""

// F3APIDefaultPort Default port for the api server
const F3APIDefaultPort = 80

// F3EnvVarsPrefix Perfix for the environmental variables
const F3EnvVarsPrefix = "f3_api"

// F3APILogFormat Default Logging format (For debug, it should be configurable)
const F3APILogFormat = "method=${method}, uri=${uri}, status=${status}\n"

func main() {
	initConfig()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: F3APILogFormat}))

	routing.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(getAddress()))
}

func initConfig() {
	viper.SetEnvPrefix(F3EnvVarsPrefix)
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", getIP(), getPort())
}

func getIP() string {
	viper.SetDefault("ip", F3APIDefaultIP)
	viper.BindEnv("ip")

	return viper.GetString("ip")
}

func getPort() string {
	viper.SetDefault("port", F3APIDefaultPort)
	viper.BindEnv("port")

	return viper.GetString("port")
}
