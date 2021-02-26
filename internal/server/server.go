package server

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Run(e *echo.Echo) {
	e.Logger.Fatal(e.Start(viper.GetString("port")))
}
