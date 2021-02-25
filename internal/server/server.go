package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Run(configPath, configName string) {
	/* Initialize viper config */
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	/* Initialize database */

	/* Initialize http server */
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(viper.GetString("port")))
}
