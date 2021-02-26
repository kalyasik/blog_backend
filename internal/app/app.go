package app

import (
	"log"

	"github.com/kalyasik/blog_backend/internal/server"
	v1 "github.com/kalyasik/blog_backend/internal/transport/http/v1"
	"github.com/kalyasik/blog_backend/pkg/database/postgres"
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
	postgres.InitDatabase()

	/* Create and run echo server */
	e := echo.New()
	e.GET("/api/v1/posts", v1.GetPostsHandler)
	e.POST("/api/v1/posts", v1.CreatePostHandler)
	e.PUT("/api/v1/posts/:id", v1.UpdatePostHandler)
	e.DELETE("/api/v1/posts/:id", v1.DeletePostHandler)
	server.Run(e)
}
