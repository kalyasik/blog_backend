package app

import (
	"log"

	_ "github.com/kalyasik/blog_backend/docs"
	"github.com/kalyasik/blog_backend/internal/server"
	v1 "github.com/kalyasik/blog_backend/internal/transport/http/v1"
	"github.com/kalyasik/blog_backend/pkg/database/postgres"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Blog API
// @version 1.0
// @description REST API server from NIX education

// @host localhost:8087
// @BasePath /api/v1

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

	/* Swagger documentation */
	e.GET("/api/v1/swagger/*", echoSwagger.WrapHandler)

	e.GET("/api/v1/posts", v1.GetPostsHandler)
	e.POST("/api/v1/posts", v1.CreatePostHandler)
	e.PUT("/api/v1/posts/:id", v1.UpdatePostHandler)
	e.DELETE("/api/v1/posts/:id", v1.DeletePostHandler)
	server.Run(e)
}
